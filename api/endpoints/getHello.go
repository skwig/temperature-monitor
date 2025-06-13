package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (e Endpoints) GetHello(c *gin.Context) {
	hello := "Hello"
	response := HelloResponse{Message: &hello}

	c.JSON(http.StatusOK, response)
}
