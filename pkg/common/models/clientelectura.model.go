package models

import (
	"time"

	"gorm.io/gorm"
)

type ClienteLectura struct {
	gorm.Model
	IDClienteContador uint
	ClienteContador   ClienteContador `gorm:"foreignKey:IDClienteContador"`
	IDEmpleado        uint
	Empleado          Empleado `gorm:"foreignKey:IDEmpleado"`
	Fecha             time.Time
	Lectura           float32 `gorm:"type:numeric(18,2)"`
	Consumo           float32 `gorm:"type:numeric(18,2)"`
	Imagen            string  `gorm:"type:varchar(128)"`
	Comentario        string  `gorm:"type:varchar(256)"`
	Estado            uint8   `gorm:"default:0"`
}
