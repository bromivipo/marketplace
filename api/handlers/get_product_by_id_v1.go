package handlers

import (
	"net/http"
	"strconv"

	"github.com/bromivipo/marketplace/api/consts"
	"github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
	"github.com/gin-gonic/gin"
)


func GetProductByIDV1(c *gin.Context) {
	idStr := c.Request.URL.Query().Get("id") 
	if idStr == "" {
		c.IndentedJSON(http.StatusBadRequest, definitions.ErrorResponse{Code: consts.BadRequest, Message: "id is missing in query"})
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, definitions.ErrorResponse{Code: consts.BadRequest, Message: "id should be convertable to int"})
		return
	}
	
	product := pgrepo.GetProductById(id)

	if product == nil {
		c.IndentedJSON(http.StatusNotFound, definitions.ErrorResponse{Code: consts.NotFound, Message: "product not found by id"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}
