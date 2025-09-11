package main

import (
    "github.com/gin-gonic/gin"
	"github.com/bromivipo/marketplace/api/handlers"
)

func main() {
    router := gin.Default()
    router.GET("/products/v1", handlers.GetProductsV1)
	router.GET("/product-by-id/v1", handlers.GetProductByIDV1)

    router.Run("localhost:8080")
}