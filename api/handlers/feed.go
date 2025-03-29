package handlers

import (
	"context"
	"fmt"
	"harry/query-overflow-feed/database"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// QueryFeedHandler is the handler for the query feed microservice

var feedCollection *mongo.Collection = database.GetCollection(database.DB, "feeds")

// Path: api\handlers\feed.go

func QueryFeedHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Get the query parameters from the request
	query := c.Query("query")

	if query == "" {
		c.JSON(400, gin.H{"error": "Query parameter is required"})
		return
	}

	// 2. Query the database for the feed data
	// Assuming you have a function to query the database
	feedData, err := feedCollection.Find(ctx, bson.M{"$text": bson.M{"$search": query}})
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error querying database: %v", err)})
		return
	}

	defer feedData.Close(ctx)

	// 3. Return the feed data as JSON response
	c.JSON(200, gin.H{"data": feedData})
}

// CreateFeedHandler is the handler for creating a new feed
func CreateFeedHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Get the feed data from the request body
	var feedData bson.M
	if err := c.ShouldBindJSON(&feedData); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err)})
		return
	}

	// 2. Insert the feed data into the database
	result, err := feedCollection.InsertOne(ctx, feedData)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error inserting data: %v", err)})
		return
	}

	// 3. Return the result as JSON response
	c.JSON(201, gin.H{"result": result})
}

// GetFeedHandler is the handler for getting a feed by ID
func GetFeedHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Get the feed ID from the URL parameters
	feedID := c.Param("id")

	// 2. Query the database for the feed data by ID
	var feedData bson.M
	err := feedCollection.FindOne(ctx, bson.M{"_id": feedID}).Decode(&feedData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{"error": "Feed not found"})
			return
		}
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error querying database: %v", err)})
		return
	}

	// 3. Return the feed data as JSON response
	c.JSON(200, gin.H{"data": feedData})
}

// UpdateFeedHandler is the handler for updating a feed by ID
func UpdateFeedHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Get the feed ID from the URL parameters
	feedID := c.Param("id")

	// 2. Get the updated feed data from the request body
	var updatedFeedData bson.M
	if err := c.ShouldBindJSON(&updatedFeedData); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err)})
		return
	}

	// 3. Update the feed data in the database
	result, err := feedCollection.UpdateOne(ctx, bson.M{"_id": feedID}, bson.M{"$set": updatedFeedData})
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error updating data: %v", err)})
		return
	}

	// 4. Return the result as JSON response
	c.JSON(200, gin.H{"result": result})
}

// DeleteFeedHandler is the handler for deleting a feed by ID
func DeleteFeedHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Get the feed ID from the URL parameters
	feedID := c.Param("id")

	// 2. Delete the feed data from the database
	result, err := feedCollection.DeleteOne(ctx, bson.M{"_id": feedID})
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error deleting data: %v", err)})
		return
	}

	// 3. Return the result as JSON response
	c.JSON(200, gin.H{"result": result})
}
