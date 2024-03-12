package models

import "gorm.io/gorm"

type Municipio struct {
	gorm.Model
	IDDepartamento uint
	Departamento   Departamento `gorm:"foreignKey:IDDepartamento"`
	Municipio      string       `gorm:"type:varchar(64);unique;not null"`
	Estado         uint8        `gorm:"default:0"`
}
