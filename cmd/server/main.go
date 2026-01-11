package main

import (
    "github.com/gin-gonic/gin"

    "gin-api/internal/albums"
)

func main() {
    r := gin.Default()

    // Initialize in-memory store and handler
    store := albums.NewStore()
    handler := albums.NewHandler(store)

    // Routes
    r.GET("/albums", handler.GetAll)
    r.GET("/albums/:id", handler.GetByID)
    r.POST("/albums", handler.Add)

    if err := r.Run("localhost:8080"); err != nil {
        panic(err)
    }
}
