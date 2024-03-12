package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CajaMov struct {
	gorm.Model
	IDCaja        uint
	Caja          Caja `gorm:"foreignKey:IDCaja"`
	Fecha         datatypes.Date
	HoraApertura  time.Time
	MontoApertura float32 `gorm:"type:numeric(18,2)"`
	HoraCierre    time.Time
	MontoCierre   float32 `gorm:"type:numeric(18,2)"`
	Estado        uint8   `gorm:"default:0"`
}
