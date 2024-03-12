package models

import "gorm.io/gorm"

type Sucursal struct {
	gorm.Model
	IDEmpresa   uint
	Empresa     Empresa `gorm:"foreignKey:IDEmpresa"`
	Codigo      string  `gorm:"type:varchar(32);unique;not null"`
	Descripcion string  `gorm:"type:varchar(128);unique;not null"`
	Direccion   string  `gorm:"type:varchar(128)"`
	Telefono    string  `gorm:"type:varchar(32)"`
	Estado      uint8   `gorm:"default:0"`
}
