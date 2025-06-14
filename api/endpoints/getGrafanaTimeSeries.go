package endpoints

import (
	"log"
	"net/http"
	"temperaturemonitor/api/sql"
	"time"

	"github.com/gin-gonic/gin"
	sf "github.com/sa-/slicefunk"
)

// GetGrafanaTimeSeries implements ServerInterface.
func (e *Endpoints) GetGrafanaTimeSeries(c *gin.Context) {
	respository, err := sql.NewDefaultSqliteRepository()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	items, err := respository.GetAll()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	response := sf.Map(items, func(item sql.SensorReading) TimeSeriesEntry {
		return TimeSeriesEntry{
			Humidity:    item.Humidity,
			Time:        time.Unix(item.ServerTimeUnix, 0),
			Session:     item.Session,
			Temperature: item.Temperature}
	})

	c.JSON(http.StatusOK, response)
}
