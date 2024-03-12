package models

import "gorm.io/gorm"

type TipoMovinv struct {
	gorm.Model
	IDEmpresa      uint
	Empresa        Empresa `gorm:"foreignKey:IDEmpresa"`
	Codigo         string  `gorm:"type:varchar(32);unique;not null"`
	Descripcion    string  `gorm:"type:varchar(128);not null"`
	SumaInventario uint8   `gorm:"default:0"`
	SumaCosto      uint8   `gorm:"default:0"`
	Estado         uint8   `gorm:"default:0"`
}
