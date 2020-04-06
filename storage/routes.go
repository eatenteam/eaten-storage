package main

import "net/http"

func (s *Server) routes() {
    s.router.HandlerFunc(http.MethodGet, "/health", s.handleHealthCheck())

    s.router.GET("/api/malls", s.handleMallsGet())
    s.router.GET("/api/malls/:id", s.extractId(s.handleMallsGetId()))
    s.router.POST("/api/malls", s.handleMallsPost())
    s.router.PUT("/api/malls/:id", s.extractId(s.handleMallsPutId()))
    s.router.DELETE("/api/malls/:id", s.extractId(s.handleMallsDeleteId()))

    s.router.GET("/api/shops", s.handleShopsGet())
    s.router.GET("/api/shops/:id", s.extractId(s.handleShopsGetId()))
    s.router.POST("/api/shops", s.handleShopsPost())
    s.router.PUT("/api/shops/:id", s.extractId(s.handleShopsPutId()))
    s.router.DELETE("/api/shops/:id", s.extractId(s.handleShopsDeleteId()))

    s.router.GET("/api/products", s.handleProductsGet())
    s.router.GET("/api/products/:id", s.extractId(s.handleProductsGetId()))
    s.router.POST("/api/products", s.handleProductsPost())
    s.router.PUT("/api/products/:id", s.extractId(s.handleProductsPutId()))
    s.router.DELETE("/api/products/:id", s.extractId(s.handleProductsDeleteId()))
}
