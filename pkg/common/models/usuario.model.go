package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	IDPerfil   uint
	Perfil     Perfil `gorm:"foreignKey:IDPerfil"`
	Usuario    string `gorm:"type:varchar(32);unique;not null"`
	Nombres    string `gorm:"type:varchar(32);not null"`
	Apellidos  string `gorm:"type:varchar(32);not null"`
	Contrasena string `gorm:"type:varchar(32);not null"`
	Estado     uint8  `gorm:"default:0"`
}
