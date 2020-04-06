package main

import "net/http"

func (s *Server) routes() {
    s.router.HandlerFunc(http.MethodGet, "/health", s.handleHealthCheck())

    s.router.GET("/api/malls", s.handleMallsGet())
    s.router.GET("/api/malls/:id", s.extractId(s.handleMallsGetId()))
    s.router.POST("/api/malls", s.handleMallsPost())
    s.router.PUT("/api/malls/:id", s.extractId(s.handleMallsPutId()))
    s.router.DELETE("/api/malls/:id", s.extractId(s.handleMallsDeleteId()))
}
