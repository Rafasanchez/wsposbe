package models

import "gorm.io/gorm"

type Departamento struct {
	gorm.Model
	Departamento string `gorm:"type:varchar(64);unique;not null"`
	IDPais       uint
	Pais         Pais  `gorm:"foreignKey:IDPais"`
	Estado       uint8 `gorm:"default:0"`
}
