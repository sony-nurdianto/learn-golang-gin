package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sony-nurdianto/go/platform/newsfeed"
)

func NewsfeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
