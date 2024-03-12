package models

import (
	"time"

	"gorm.io/gorm"
)

type Movinvh struct {
	gorm.Model
	IDEmpresa           uint
	Empresa             Empresa `gorm:"foreignKey:IDEmpresa"`
	IDProveedor         uint
	Proveedor           Proveedor `gorm:"foreignKey:IDProveedor"`
	IDTipomovinv        uint
	TipoMovinv          TipoMovinv `gorm:"foreignKey:IDTipomovinv"`
	Fecha               time.Time
	SerieReferencia     string `gorm:"type:varchar(32)"`
	DocumentoReferencia string `gorm:"type:varchar(32)"`
	Descripcion         string `gorm:"type:varchar(128)"`
	Estado              uint8  `gorm:"default:0"`
}
