package controllers

import (
	"fmt"

	"github.com/SFerguson90/FergusonAero/inits"
	"github.com/SFerguson90/FergusonAero/models"

	"github.com/gin-gonic/gin"
)

func CreateAirplane(ctx *gin.Context) {

	var body struct {
		Manufacturer string
		Airframe     string
		Series       string
		TailNumber   string
		Hobbs        uint
		Tachometer   uint
		IsOperable   bool
	}

	ctx.BindJSON(&body)

	airplane := models.Airplane{
		Manufacturer: body.Manufacturer,
		Airframe:     body.Airframe,
		Series:       body.Series,
		TailNumber:   body.TailNumber,
		Hobbs:        body.Hobbs,
		Tachometer:   body.Tachometer,
		IsOperable:   body.IsOperable,
	}

	fmt.Println(airplane)

	result := inits.DB.Create(&airplane)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": airplane})
}

func GetAirplanes(ctx *gin.Context) {

	var airplanes []models.Airplane

	result := inits.DB.Find(&airplanes)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": airplanes})
}

func GetAirplane(ctx *gin.Context) {

	var airplane models.Airplane

	result := inits.DB.First(&airplane, ctx.Param("id"))

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": airplane})
}

func UpdateAirplane(ctx *gin.Context) {

	var body struct {
		Manufacturer string
		Airframe     string
		Series       string
		TailNumber   string
		Hobbs        uint
		Tachometer   uint
		IsOperable   bool
	}

	ctx.BindJSON(&body)

	var airplane models.Airplane

	result := inits.DB.First(&airplane, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&airplane).Updates(models.Airplane{
		Manufacturer: body.Manufacturer,
		Airframe:     body.Airframe,
		Series:       body.Series,
		TailNumber:   body.TailNumber,
		Hobbs:        body.Hobbs,
		Tachometer:   body.Tachometer,
		IsOperable:   body.IsOperable,
	})

	ctx.JSON(200, gin.H{"data": airplane})

}

func DeleteAirplane(ctx *gin.Context) {

	id := ctx.Param("id")

	inits.DB.Delete(&models.Airplane{}, id)

	ctx.JSON(200, gin.H{"data": "Airplane has been deleted successfully"})
}
