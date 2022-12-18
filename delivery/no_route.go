package delivery

import (
	"fmt"
	"net/http"
	"service-campaign-startup/model/dto"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	response := dto.BuildResponse(
		"User not found",
		"FAILED",
		http.StatusNotFound,
		map[string]interface{}{"ERROR": fmt.Sprintf("Resource or route %s is not found", c.Request.URL.Path)},
	)
	c.AbortWithStatusJSON(http.StatusNotFound, response)
}
