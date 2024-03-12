package models

import (
	"time"

	"gorm.io/gorm"
)

type ClienteContador struct {
	gorm.Model
	IDClienteDireccion uint
	ClienteDireccion   ClienteDireccion `gorm:"foreignKey:IDClienteDireccion"`
	Codigo             string           `gorm:"type:varchar(32);unique;not null"`
	FechaAlta          time.Time
	Estado             uint8 `gorm:"default:0"`
}
