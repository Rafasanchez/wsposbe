package models

import (
	"time"

	"gorm.io/gorm"
)

type Ordenh struct {
	gorm.Model
	IDSucursal  uint
	Sucursal    Sucursal `gorm:"foreignKey:IDSucursal"`
	IDCliente   uint
	Cliente     Cliente `gorm:"foreignKey:IDCliente"`
	IDFormaPago uint
	FormaPago   FormaPago `gorm:"foreignKey:IDFormaPago"`
	FechaInicio time.Time
	NoPlaca     string  `gorm:"type:varchar(8)"`
	PagaCon     float32 `gorm:"type:numeric(18,2)"`
	Total       float32 `gorm:"type:numeric(18,2)"`
	FechaFin    time.Time
	Estado      uint8 `gorm:"default:0"`
}
