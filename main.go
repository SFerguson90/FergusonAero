package main

import (
	"HubCityAviation/middleware"
	"HubCityAviation/routes/airplanes"
	"HubCityAviation/routes/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Setup middleware
	r.Use(middleware.KerberosAuth())

	// Register routes
	users.RegisterRoutes(r)
	airplanes.RegisterRoutes(r)
	// ... register other route groups as needed

	r.Run() // Start the server
}
