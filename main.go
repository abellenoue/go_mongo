package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"cesi/go_mongo/config"
	"cesi/go_mongo/routes"
)

func main() {
	// Database
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
