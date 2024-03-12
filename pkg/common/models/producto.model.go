package models

import "gorm.io/gorm"

type Producto struct {
	gorm.Model
	IDEmpresa    uint
	Empresa      Empresa `gorm:"foreignKey:IDEmpresa"`
	Codigo       string  `gorm:"type:varchar(32);unique;not null"`
	Descripcion  string  `gorm:"type:varchar(128);not null"`
	Imagen       string  `gorm:"type:varchar(256)"`
	CobroxTiempo uint8   `gorm:"default:0"`
	GeneraTicket uint8   `gorm:"default:0"`
	CargoParqueo uint8   `gorm:"default:0"`
	Estado       uint8   `gorm:"default:0"`
}
