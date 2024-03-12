package models

import "gorm.io/gorm"

type Ordend struct {
	gorm.Model
	IDOrdenh    uint
	Ordenh      Ordenh `gorm:"foreignKey:IDOrdenh"`
	IDProducto  uint
	Producto    Producto `gorm:"foreignKey:IDProducto"`
	Codigo      string   `gorm:"type:varchar(32)"`
	Descripcion string   `gorm:"type:varchar(128)"`
	Cantidad    float32  `gorm:"type:numeric(18,2)"`
	Precio      float32  `gorm:"type:numeric(18,2)"`
	SubTotal    float32  `gorm:"type:numeric(18,2)"`
}
