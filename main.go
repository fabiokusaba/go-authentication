package main

import (
	"authentication/config"
	"authentication/helpers"
	"authentication/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting application...")

	key := config.GenerateRandomKey()
	helpers.SetJWTKey(key)
	fmt.Printf("Generated key: %s\n", key)

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
	log.Println("Server is running on port: 8080")
}
