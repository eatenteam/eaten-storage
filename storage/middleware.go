package main

import (
    "net/http"
    "context"

    httprouter  "github.com/julienschmidt/httprouter"
    primitive   "go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) extractId(fn httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
        if hexId := params.ByName("id"); len(hexId) > 0 {
            id, err := primitive.ObjectIDFromHex(hexId)
            if err != nil {
                respondErr(w, r, http.StatusBadRequest, "failed to validate the specified ID")
                return
            }
            ctx := context.WithValue(r.Context(), "id", id)
            fn(w, r.WithContext(ctx), params)
            return
        }
        fn(w, r, params)
    }
}
