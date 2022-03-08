package handler

import (
	"gin-RESTApi/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNewsfeed(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
