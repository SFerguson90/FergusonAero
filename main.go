package main

import (
	"github.com/SFerguson90/HubCityAviation/controllers"
	"github.com/SFerguson90/HubCityAviation/inits"
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

	r.Run()
}
