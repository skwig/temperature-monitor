package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGrafanaTimeSeries implements ServerInterface.
func (e *Endpoints) GetGrafanaTimeSeries(c *gin.Context) {
	response := make([]TimeSeriesEntry, 1)

	c.JSON(http.StatusOK, response)
}
