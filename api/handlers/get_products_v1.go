package handlers

import (
	"net/http"

	"github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
	"github.com/gin-gonic/gin"
)

func GetProductsV1(c *gin.Context) {
	products := definitions.GetProductsV1Response{Products: pgrepo.GetProducts()}
    c.IndentedJSON(http.StatusOK, products)
}
