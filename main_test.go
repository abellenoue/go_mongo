package main

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"

	"cesi/go_mongo/config"
	"cesi/go_mongo/routes"
)

func TestMain(m *testing.M) {
	// Database
	config.TestConnect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)
	// controllers.TestCreateTodo(m)
	m.Run()

	log.Fatal(router.Run(":493"))
}
