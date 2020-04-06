package main

import (
    "context"
    "net/http"
    "log"
    "time"

    envdecode   "github.com/joeshaw/envdecode"
    httprouter  "github.com/julienschmidt/httprouter"
    mongo       "go.mongodb.org/mongo-driver/mongo"
    options     "go.mongodb.org/mongo-driver/mongo/options"
    readpref    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
    var addr, mongo string
    if err := loadEnv(&mongo, &addr); err != nil {
        log.Fatalln("Loading Environment Error: ", err)
    }
    db, err := connectMongo(mongo)
    if err != nil {
        log.Fatalln("Mongo Error: ", err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    defer db.Disconnect(ctx)
    router, err := createRouter()
    if err != nil {
        log.Fatalln("HTTPRouter Error: ", err)
    }
    s := newServer(db, router)
    log.Println("🌏  Server listening on ", addr)
    log.Fatalln(http.ListenAndServe(addr, s))
}

func loadEnv(mongo, addr *string) error {
    var ts struct {
        Mongo   string  `env:"MONGO,required"`
        Port    string  `env:"PORT,required"`
    }
    if err := envdecode.Decode(&ts); err != nil {
        return err
    }
    *mongo = ts.Mongo
    *addr = ":" + ts.Port
    return nil
}

func connectMongo(uri string) (*mongo.Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    log.Println("🔥  Connecting to MongoDB Cluster...")
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }
    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    log.Println("🔥  Pinging to MongoDB Cluster...")
    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        return nil, err
    }
    log.Println("✅  Successfully connect MongoDB Cluster")
    return client, nil
}

func createRouter() (*httprouter.Router, error) {
    router := httprouter.New()
    return router, nil
}
