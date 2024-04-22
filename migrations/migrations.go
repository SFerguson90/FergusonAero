package main

import (
	"github.com/SFerguson90/HubCityAviation/inits"
	"github.com/SFerguson90/HubCityAviation/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Airplane{})
}
