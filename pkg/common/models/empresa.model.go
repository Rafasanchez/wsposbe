package models

import "gorm.io/gorm"

type Empresa struct {
	gorm.Model
	Empresa   string `gorm:"type:varchar(32);unique;not null"`
	Direccion string `gorm:"type:varchar(128)"`
	Nit       string `gorm:"type:varchar(32)"`
	Telefono  string `gorm:"type:varchar(32)"`
	IDPais    uint
	Pais      Pais  `gorm:"foreignKey:IDPais"`
	Estado    uint8 `gorm:"default:0"`
}
