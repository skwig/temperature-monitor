package endpoints

import (
	"log"
	"net/http"

	"temperaturemonitor/api/sql"

	"github.com/gin-gonic/gin"
)

// IngestFromSensor implements ServerInterface.
func (e *Endpoints) IngestFromSensor(c *gin.Context) {
	var requestBody IngestFromSensorJSONRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.Status(http.StatusBadRequest)
	}

	respository, err := sql.NewDefaultSqliteRepository()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v\n", requestBody)

	reading := sql.SensorReading{
		Session:        requestBody.Session,
		SensorTimeUnix: requestBody.SensorTime.Unix(),
		Temperature:    requestBody.Temperature,
		Humidity:       requestBody.Humidity}

	err = respository.Save(&reading)
	if err != nil {
		log.Println(err)
	}

	c.Status(http.StatusOK)
}
