package models

import (
	"time"

	"gorm.io/gorm"
)

type Empleado struct {
	gorm.Model
	IDSucursal        uint
	Sucursal          Sucursal `gorm:"foreignKey:IDSucursal"`
	Codigo            string   `gorm:"type:varchar(32);unique;not null"`
	PrimerNombre      string   `gorm:"type:varchar(32);not null"`
	SegundoNombre     string   `gorm:"type:varchar(32)"`
	PrimerApellido    string   `gorm:"type:varchar(32);not null"`
	SegundoApellido   string   `gorm:"type:varchar(32)"`
	ApellidoCasada    string   `gorm:"type:varchar(32)"`
	CorreoElectronico string   `gorm:"type:varchar(64)"`
	FechaNacimiento   time.Time
	Estado            uint8 `gorm:"default:0"`
}
