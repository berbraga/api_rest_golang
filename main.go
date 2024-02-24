package main

import (
	"github.com/gin-gonic/gin"
	"numismatico/api/routes"
)

func main() {
	app := gin.Default()
	routes.AppRoutes(app)
	app.Run("localhost:6789")
}
