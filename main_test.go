package main

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"

	"cesi/go_mongo/config"
	"cesi/go_mongo/routes"
)

func Testmain(m *testing.M) {
	// Database
	config.TestConnect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	m.Run()

	log.Fatal(router.Run(":4747"))
}
