package endpoints

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IngestFromSensor implements ServerInterface.
func (e *Endpoints) IngestFromSensor(c *gin.Context) {
	var requestBody IngestFromSensorJSONRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.Status(http.StatusBadRequest)
	}

	log.Printf("%+v\n", requestBody)

	c.Status(http.StatusOK)
}
