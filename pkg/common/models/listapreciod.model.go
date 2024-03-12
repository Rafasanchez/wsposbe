package models

import "gorm.io/gorm"

type Listapreciod struct {
	gorm.Model
	IDListaprecioh uint
	Listaprecioh   Listaprecioh `gorm:"foreignKey:IDListaprecioh"`
	IDProducto     uint
	Producto       Producto `gorm:"foreignKey:IDProducto"`
	Precio         float32  `gorm:"type:numeric(18,2)"`
}
