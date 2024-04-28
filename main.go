package main

import (
	"os"

	"github.com/SFerguson90/FergusonAero/controllers"
	"github.com/SFerguson90/FergusonAero/inits"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.GET("/airplanes", controllers.GetAirplanes)
	r.POST("/airplanes", controllers.CreateAirplane)

	r.GET("/airplane/:id", controllers.GetAirplane)
	r.PUT("/airplane/:id", controllers.UpdateAirplane)
	r.DELETE("/airplane/:id", controllers.DeleteAirplane)

	port := os.Getenv("HTTP_PLATFORM_PORT")

	if port == "" {
		port = "8080"
	}

	r.Run("127.0.0.1:" + port)
}
