package main

import (
    "context"
    "net/http"

    httprouter  "github.com/julienschmidt/httprouter"
    bson        "go.mongodb.org/mongo-driver/bson"
    primitive   "go.mongodb.org/mongo-driver/bson/primitive"
    mongo       "go.mongodb.org/mongo-driver/mongo"
    options     "go.mongodb.org/mongo-driver/mongo/options"
)

type Shop struct {
    Id          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
    Brand       string              `bson:"brand" json:"brand"`
    Tel         string              `bson:"tel" json:"tel"`
    Location    Location            `bson:"location" json:"location"`
    Stock       []Stock             `bson:"stock,omitempty" json:"stock"`
    Open        string              `bson:"open" json:"open"`
    Close       string              `bson:"close" json:"close"`
}

type Stock struct {
    Product     primitive.ObjectID  `bson:"product" json:"product"`
    Quantity    int                 `bson:"quantity" json:"quantity"`
}

type Location struct {
    Mall        primitive.ObjectID  `bson:"mall,omitempty" json:"mall"`
    Address     string              `bson:"addr,omitempty json:"addr"`
}

func (s *Server) handleShopsGet() httprouter.Handle {
    collection := s.db.Database("storage").Collection("shops")
    findOptions := options.Find()
    filter := bson.M{}
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var results []*Shop
        cur, err := collection.Find(context.TODO(), filter, findOptions)
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, "failed to find documents")
            return
        }
        defer cur.Close(context.TODO())
        for cur.Next(context.TODO()) {
            var elem Shop
            err := cur.Decode(&elem)
            if err != nil {
                respondErr(w, r, http.StatusInternalServerError, err)
                return
            }
            results = append(results, &elem)
        }
        if err := cur.Err(); err != nil {
            respondErr(w, r, http.StatusInternalServerError, err)
        }
        respond(w, r, http.StatusOK, &results)
    }
}

func (s *Server) handleShopsGetId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("shops")
    findOptions := options.FindOne()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var result Shop
        filter := bson.M{"_id": r.Context().Value("id")}
        err := collection.FindOne(context.TODO(), filter, findOptions).Decode(&result)
        if err != nil {
            if err == mongo.ErrNoDocuments {
                respondErr(w, r, http.StatusNotFound, err)
                return
            }
            respondErr(w, r, http.StatusInternalServerError, "failed to find document from the database")
            return
        }
        respond(w, r, http.StatusOK, &result)
    }
}

func (s *Server) handleShopsPost() httprouter.Handle {
    collection := s.db.Database("storage").Collection("shops")
    insertOptions := options.InsertOne()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var shop Shop
        if err := decodeBody(r, &shop); err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to read shop from request")
            return
        }
        result, err := collection.InsertOne(context.TODO(), shop, insertOptions)
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, "failed to insert document into the database")
            return
        }
        respond(w, r, http.StatusCreated, &result)
    }
}

func (s *Server) handleShopsPutId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("shops")
    updateOptions := options.FindOneAndReplace()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var shop, result Shop
        filter := bson.M{"_id": r.Context().Value("id")}
        if err := decodeBody(r, &shop); err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to read shop from request")
            return
        }
        err := collection.FindOneAndReplace(context.TODO(), filter, shop, updateOptions).Decode(&result)
        if err != nil {
            if err == mongo.ErrNoDocuments {
                respondErr(w, r, http.StatusNotFound, err)
                return
            }
            respondErr(w, r, http.StatusInternalServerError, "failed to update a document in the database")
            return
        }
        respond(w, r, http.StatusOK, &result)
    }
}

func (s *Server) handleShopsDeleteId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("shops")
    deleteOptions := options.FindOneAndDelete()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var result Shop
        filter := bson.M{"_id": r.Context().Value("id")}
        err := collection.FindOneAndDelete(context.TODO(), filter, deleteOptions).Decode(&result)
        if err != nil {
            if err == mongo.ErrNoDocuments {
                respondErr(w, r, http.StatusNotFound, err)
                return
            }
            respondErr(w, r, http.StatusInternalServerError, "failed to delete a document from the database")
            return
        }
        respond(w, r, http.StatusOK, &result)
    }
}
