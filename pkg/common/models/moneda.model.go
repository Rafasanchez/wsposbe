package models

import (
	"gorm.io/gorm"
)

type Moneda struct {
	gorm.Model
	Moneda       string `gorm:"type:varchar(64);unique;not null"`
	CodigoMoneda string `gorm:"type:varchar(16);unique;not null"`
	Estado       uint8  `gorm:"default:0"`
}
