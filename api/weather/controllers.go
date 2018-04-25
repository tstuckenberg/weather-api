package weather

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//---- Allocation Request API ----

//ConditionsResponse type
type ConditionsResponse struct {
	Weather Conditions `json:"weather"`
	Success bool       `json:"success"`
}

//ErrorResponse type
type ErrorResponse struct {
	Message error `json:"message,omitempty"`
	Success bool  `json:"success"`
}

//RenderError function
func (r *Resource) RenderError(c *gin.Context, status int, message error) {
	c.JSON(status, ErrorResponse{
		Message: message,
		Success: false,
	})
}

//HandleGetWeather decommissions an Allocation
func (r *Resource) HandleGetWeather(c *gin.Context) {

	if c.Request.URL.Query()["city"] == nil {

		// if len(c.Request.URL.Query()["city"]) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: errors.New("Please specify a location"),
			Success: false,
		})
	} else {

		city := c.Request.URL.Query()["city"][0]

		weather, err := GetWeather(city)

		if err != nil {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Message: err,
				Success: false,
			})
		} else {
			c.JSON(http.StatusOK, ConditionsResponse{
				Weather: weather,
				Success: true,
			})
		}
	}
}
