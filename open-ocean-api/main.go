package main

import (
    "github.com/gin-gonic/gin"
    "open-ocean-api/models"
)

func main() {
    r := gin.Default()
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, "ok!")
    })
    models.SetUpModels()
    r.Run()
}
