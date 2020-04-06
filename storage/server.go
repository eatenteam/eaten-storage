package main

import (
    "net/http"

    mongo       "go.mongodb.org/mongo-driver/mongo"
    httprouter  "github.com/julienschmidt/httprouter"
)

type Server struct {
    db      *mongo.Client
    router  *httprouter.Router
}

func newServer(db *mongo.Client, router *httprouter.Router) *Server {
    s := &Server{db, router}
    s.routes()
    return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.router.ServeHTTP(w, r)
}
