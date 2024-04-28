package main

import (
	"github.com/SFerguson90/FergusonAero/inits"
	"github.com/SFerguson90/FergusonAero/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Airplane{})
}
