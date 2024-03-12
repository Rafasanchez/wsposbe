package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Menu        string `gorm:"type:varchar(32);unique;not null"`
	Descripcion string `gorm:"type:varchar(128);not null"`
	Imagen      string `gorm:"type:varchar(128)"`
	Padre       string `gorm:"type:varchar(32)"`
	IDProducto  uint
	Producto    Producto `gorm:"foreignKey:IDProducto"`
	Estado      uint8    `gorm:"default:0"`
}
