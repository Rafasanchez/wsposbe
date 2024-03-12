package models

import "gorm.io/gorm"

type Proveedor struct {
	gorm.Model
	IDEmpresa         uint
	Empresa           Empresa `gorm:"foreignKey:IDEmpresa"`
	Nombres           string  `gorm:"type:varchar(32);not null"`
	Apellidos         string  `gorm:"type:varchar(32);not null"`
	CorreoElectronico string  `gorm:"type:varchar(64)"`
	Nit               string  `gorm:"type:varchar(32);unique;not null"`
	DPI               string  `gorm:"type:varchar(32);unique;not null"`
	Direccion         string  `gorm:"type:varchar(256)"`
	Telefono          string  `gorm:"type:varchar(32)"`
	Estado            uint8   `gorm:"default:0"`
}
