//  creating server for the query feed mircoservice
//  will be using gin framework here to deinf the route that will be used to query the feed handler

// Path: cmd\server.go

package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"harry/query-overflow-feed/api/handlers"
)

func NewServer() *gin.Engine {
	fmt.Println("Starting server...")

	router := gin.Default()

	// TODO: DB connection

	//  Define the route for the query feed handler
	feeds := router.Group("/feeds")
	{
		feeds.GET("/query", handlers.QueryFeedHandler)
	}
	return router
}
