package endpoints

import (
	"log"
	"net/http"
	"time"

	"temperaturemonitor/api/sql"

	"github.com/gin-gonic/gin"
)

// IngestFromSensor implements ServerInterface.
func (e *Endpoints) IngestFromSensor(c *gin.Context) {
	var requestBody IngestFromSensorJSONRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	respository, err := sql.NewDefaultSqliteRepository()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	reading := sql.SensorReading{
		Session:        requestBody.Session,
		ServerTimeUnix: time.Now().UTC().Unix(),
		SensorTimeUnix: requestBody.SensorTime.Unix(),
		Temperature:    requestBody.Temperature,
		Humidity:       requestBody.Humidity}

	err = respository.Save(&reading)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
