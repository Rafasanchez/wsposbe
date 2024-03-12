package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Caja struct {
	gorm.Model
	IDEmpresa         uint
	Empresa           Empresa `gorm:"foreignKey:IDEmpresa"`
	TipoAutorizacion  string  `gorm:"type:varchar(32);not null"`
	FechaAutorizacion datatypes.Date
	Resolucion        string `gorm:"type:varchar(64);unique;not null"`
	Serie             string `gorm:"type:varchar(64);not null"`
	InicioCorrelativo uint   `gorm:"default:0"`
	FinalCorrelativo  uint   `gorm:"default:0"`
	Estado            uint8  `gorm:"default:0"`
}
