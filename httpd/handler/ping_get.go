package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Handler test
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "FoundMe",
		})
	}
}
