package main

import (
	"go-backend/config"
	"go-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.AuthRoutes(r)

	r.Run(":8080")

}

// go mod init go-backend
// go mod tidy
// go run cmd/main.go
