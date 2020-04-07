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

type Mall struct {
    Id          primitive.ObjectID      `bson:"_id,omitempty" json:"id"`
    Brand       string                  `bson:"brand" json:"brand"`
    Description string                  `bson:"description,omitempty" json:"description"`
    Shops       []primitive.ObjectID    `bson:"shops,omitempty"json:"shops"`
}

func (s *Server) handleMallsGet() httprouter.Handle {
    collection := s.db.Database("storage").Collection("malls")
    findOptions := options.Find()
    filter := bson.M{}
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var results []*Mall
        cur, err := collection.Find(context.TODO(), filter, findOptions)
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, err)
            return
        }
        defer cur.Close(context.TODO())
        for cur.Next(context.TODO()) {
            var elem Mall
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

func (s *Server) handleMallsGetId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("malls")
    findOptions := options.FindOne()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var result Mall
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

func (s *Server) handleMallsPost() httprouter.Handle {
    collection := s.db.Database("storage").Collection("malls")
    insertOptions := options.InsertOne()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var mall Mall
        if err := decodeBody(r, &mall); err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to read mall from request")
            return
        }
        result, err := collection.InsertOne(context.TODO(), mall, insertOptions)
        if err != nil {
            respondErr(w, r, http.StatusInternalServerError, "failed to insert document into the database")
            return
        }
        respond(w, r, http.StatusCreated, &result)
    }
}

func (s *Server) handleMallsPutId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("malls")
    updateOptions := options.FindOneAndReplace()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var mall, result Mall
        filter := bson.M{"_id": r.Context().Value("id")}
        if err := decodeBody(r, &mall); err != nil {
            respondErr(w, r, http.StatusBadRequest, "failed to read mall from request")
            return
        }
        err := collection.FindOneAndReplace(context.TODO(), filter, mall, updateOptions).Decode(&result)
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

func (s *Server) handleMallsDeleteId() httprouter.Handle {
    collection := s.db.Database("storage").Collection("malls")
    deleteOptions := options.FindOneAndDelete()
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        var result Mall
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
