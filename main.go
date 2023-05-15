// go query feed microservice

package main

import (
	"fmt"
	"log"

	"os"

	server "harry/query-overflow-feed/cmd"
)

func main() {
	fmt.Println("Reading config...")

	// reading the port from the environment variable
	if port := os.Getenv("PORT"); port != "" {
		port = "8080"
	}

	// creating the server
	server := server.NewServer()

	// running the server
	log.Printf("Listening on port %s", "8080")
	err := server.Run(":8080")

	// checking if there is any error while running the server
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
