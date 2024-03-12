package models

import (
	"time"

	"gorm.io/gorm"
)

type Listaprecioh struct {
	gorm.Model
	IDEmpresa   uint
	Empresa     Empresa `gorm:"foreignKey:IDEmpresa"`
	Codigo      string  `gorm:"type:varchar(32);unique;not null"`
	Descripcion string  `gorm:"type:varchar(128);not null"`
	FechaDesde  time.Time
	FechaHasta  time.Time
	Estado      uint8 `gorm:"default:0"`
}
