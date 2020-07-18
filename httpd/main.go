package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sony-nurdianto/go/httpd/handler"
	"github.com/sony-nurdianto/go/platform/newsfeed"
)

func main() {

	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/PostN", handler.NewsfeedPost(feed))

	r.Run()

	// feed := newsfeed.New()

	// fmt.Println(feed)

	// feed.Add(newsfeed.Item{"This is Title", "This is Post", "This is Author"})

	// fmt.Println(feed)
}
