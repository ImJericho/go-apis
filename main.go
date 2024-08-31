// File: main.go
package main

import (
	"apis/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterUserRoutes(router)
	router.Run(":8080")
}
