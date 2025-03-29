//  creating server for the query feed mircoservice
//  will be using gin framework here to deinf the route that will be used to query the feed handler

// Path: cmd\server.go

package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"harry/query-overflow-feed/api/handlers"
	"harry/query-overflow-feed/utils"
)

func NewServer() *gin.Engine {
	fmt.Println("Starting server...")

	router := gin.Default()

	// Define the route for the health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	//  Define the route for the query feed handler
	feeds := router.Group("/feeds")
	{
		// Middleware to validate the token for all routes in this group
		feeds.Use(utils.ValidateToken)

		// Define the routes for the feeds
		feeds.GET("/query", handlers.QueryFeedHandler)
		feeds.POST("/create", handlers.CreateFeedHandler)
		feeds.GET("/get/:id", handlers.GetFeedHandler)
		feeds.PUT("/update/:id", handlers.UpdateFeedHandler)
		feeds.DELETE("/delete/:id", handlers.DeleteFeedHandler)

	}

	return router
}
