package models

import "gorm.io/gorm"

type Perfil struct {
	gorm.Model
	Codigo      string `gorm:"type:varchar(32);unique;not null"`
	Descripcion string `gorm:"type:varchar(128);not null"`
	Estado      uint8  `gorm:"default:0"`
}
