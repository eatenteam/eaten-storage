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

type Product struct {
    Id          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
    Name        string              `bson:"name" json:"name"`
    Price       int                 `bson:"price" json:"price"`
}

func (s *Server) handleProductsGet() httprouter.Handle {
    collection := s.db.Database("storage").Collection("products")
    findOptions := options.Find()
    filter := bson.M{}
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var results []*Product
        cur, err := collection.Find(context.TODO(), filter, findOptions)
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, err)
            return
        }
        defer cur.Close(context.TODO())
        for cur.Next(context.TODO()) {
            var elem Product
            err := cur.Decode(&elem)
            if err != nil {
                respondErr(w, r, http.StatusInternalServerError, err)
                return
            }
            results = append(results, &elem)
        }
        if err := cur.Err(); err != nil {
            respondErr(w, r, http.StatusInternalServerError, err)
            return
        }
        respond(w, r, http.StatusOK, &results)
    }
}

func (s *Server) handleProductsGetId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("products")
    findOptions := options.FindOne()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var result Product
        filter := bson.M{"_id": r.Context().Value("id")}
        err := collection.FindOne(context.TODO(), filter, findOptions).Decode(&result)
        if err != nil {
            if err == mongo.ErrNoDocuments {
                respondErr(w, r, http.StatusNotFound, err)
                return
            }
            respondErr(w, r, http.StatusInternalServerError, "failed to find a document from the database")
            return
        }
        respond(w, r, http.StatusOK, &result)
    }
}

func (s *Server) handleProductsPost() httprouter.Handle {
    collection := s.db.Database("storage").Collection("products")
    insertOptions := options.InsertOne()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var product Product
        if err := decodeBody(r, &product); err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to read product from request")
            return
        }
        result, err := collection.InsertOne(context.TODO(), product, insertOptions)
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, "failed to insert document into the database")
            return
        }
        respond(w, r, http.StatusCreated, &result)
    }
}

func (s *Server) handleProductsPutId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("products")
    updateOptions := options.FindOneAndReplace()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var product, result Product
        filter := bson.M{"_id": r.Context().Value("id")}
        if err := decodeBody(r, &product); err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to read product from request")
            return
        }
        err := collection.FindOneAndReplace(context.TODO(), filter, product, updateOptions).Decode(&result)
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

func (s *Server) handleProductsDeleteId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("products")
    deleteOptions := options.FindOneAndDelete()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var result Product
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
