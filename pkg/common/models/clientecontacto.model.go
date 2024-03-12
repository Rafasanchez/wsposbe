package models

import "gorm.io/gorm"

type ClienteContacto struct {
	gorm.Model
	IDCliente      uint
	Cliente        Cliente `gorm:"foreignKey:IDCliente"`
	IDTipoContacto uint
	TipoContacto   TipoContacto `gorm:"foreignKey:IDTipoContacto"`
	Contacto       string       `gorm:"type:varchar(32);unique;not null"`
	Estado         uint8        `gorm:"default:0"`
}
