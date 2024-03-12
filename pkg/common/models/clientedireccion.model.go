package models

import (
	"gorm.io/gorm"
)

type ClienteDireccion struct {
	gorm.Model
	IDMunicipio uint
	Municipio   Municipio `gorm:"foreignKey:IDMunicipio"`
	IDCliente   uint
	Cliente     Cliente `gorm:"foreignKey:IDCliente"`
	Direccion   string  `gorm:"type:varchar(512);not null"`
	Latitud     float32 `gorm:"type:numeric(12,9)"`
	Longitud    float32 `gorm:"type:numeric(12,9)"`
	Estado      uint8   `gorm:"default:0"`
}
