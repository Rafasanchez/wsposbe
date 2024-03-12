package models

import (
	"gorm.io/gorm"
)

type Pais struct {
	gorm.Model
	Pais       string `gorm:"type:varchar(64);unique;not null"`
	CodigoPais string `gorm:"type:varchar(16);unique;not null"`
	IDMoneda   uint
	Moneda     Moneda `gorm:"foreignKey:IDMoneda"`
	Estado     uint8  `gorm:"default:0"`
}
