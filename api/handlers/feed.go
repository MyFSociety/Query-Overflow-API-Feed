package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// QueryFeedHandler is the handler for the query feed microservice

// Path: api\handlers\feed.go

func QueryFeedHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Query feed handler",
	})
	fmt.Println("Query feed handler")
}
