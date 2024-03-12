package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	IDEmpresa         uint
	Empresa           Empresa `gorm:"foreignKey:IDEmpresa"`
	PrimerNombre      string  `gorm:"type:varchar(32);not null"`
	SegundoNombre     string  `gorm:"type:varchar(32)"`
	PrimerApellido    string  `gorm:"type:varchar(32);not null"`
	SegundoApellido   string  `gorm:"type:varchar(32)"`
	ApellidoCasada    string  `gorm:"type:varchar(32)"`
	CorreoElectronico string  `gorm:"type:varchar(64)"`
	FechaNacimiento   time.Time
	Nit               string `gorm:"type:varchar(32);unique;not null"`
	DPI               string `gorm:"type:varchar(32);unique;not null"`
	Telefono          string `gorm:"type:varchar(32)"`
	Estado            uint8  `gorm:"default:0"`
}
