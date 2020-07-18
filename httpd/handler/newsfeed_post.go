package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sony-nurdianto/go/platform/newsfeed"
)

type newsfeedPost struct {
	Title  string `json:"title"`
	Post   string `json:"post"`
	Author string `json:"author"`
}

func NewsfeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := newsfeedPost{}
		c.Bind(&requestBody)

		item := newsfeed.Item{

			Title:  requestBody.Title,
			Post:   requestBody.Post,
			Author: requestBody.Author,
		}

		feed.Add(item)

		c.JSON(http.StatusOK, requestBody)

		fmt.Println(item)
		fmt.Println(requestBody)
		fmt.Println(feed)
	}
}
