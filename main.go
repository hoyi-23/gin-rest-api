package main

import (
	"gin-RESTApi/httpd/handler"
	"gin-RESTApi/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	feed := newsfeed.New()

	r := gin.Default()
	r.GET("/ping", handler.GetPing())
	r.GET("/newsfeed", handler.GetNewsfeed(feed))
	r.POST("/newsfeed", handler.PostNewsfeed(feed))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
