package models

import "gorm.io/gorm"

type Sysparam struct {
	gorm.Model
	Parametro   string `gorm:"type:varchar(32);unique;not null" binding:"required"`
	Descripcion string `gorm:"type:varchar(128);not null"`
	Valor       string `gorm:"type:varchar(128);not null" binding:"required"`
	Tipo        string `gorm:"type:varchar(128);not null"`
	Estado      uint8  `gorm:"default:0"`
}
