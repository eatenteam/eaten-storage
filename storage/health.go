package main

import (
    "net/http"
    "log"
)

func (s *Server) handleHealthCheck() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        respond(w, r, http.StatusOK, nil)
        log.Println("🌡️  Healthcheck passed")
        return
    }
}
