package models

import "gorm.io/gorm"

type Movinvd struct {
	gorm.Model
	IDMovinvh     uint
	Movinvh       Movinvh `gorm:"foreignKey:IDMovinvh"`
	IDProducto    uint
	Producto      Producto `gorm:"foreignKey:IDProducto"`
	Cantidad      float32  `gorm:"type:numeric(18,2)"`
	CostoUnitario float32  `gorm:"type:numeric(18,2)"`
	CostoTotal    float32  `gorm:"type:numeric(18,2)"`
}
