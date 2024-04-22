package models

import (
	"gorm.io/gorm"
)

type Airplane struct {
	gorm.Model
	Manufacturer string
	Airframe     string
	Series       string
	TailNumber   string
	Hobbs        uint
	Tachometer   uint
	IsOperable   bool
}
