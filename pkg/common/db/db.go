package db

import (
	"log"
	"wsposbe/pkg/common/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			//NoLowerCase:   true,
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	//db.Migrator().DropTable(&models.CajaMov{})
	//db.Migrator().DropTable(&models.Caja{})

	//Rutina para eliminar todas las tablas, esto se debe quitar

	db.Migrator().DropTable(&models.Menu{})
	db.Migrator().DropTable(&models.Ordend{})
	db.Migrator().DropTable(&models.Ordenh{})
	db.Migrator().DropTable(&models.Usuario{})
	db.Migrator().DropTable(&models.Perfil{})
	db.Migrator().DropTable(&models.CajaMov{})
	db.Migrator().DropTable(&models.Caja{})
	db.Migrator().DropTable(&models.ClienteLectura{})
	db.Migrator().DropTable(&models.ClienteContador{})
	db.Migrator().DropTable(&models.Movinvd{})
	db.Migrator().DropTable(&models.Movinvh{})
	db.Migrator().DropTable(&models.Listapreciod{})
	db.Migrator().DropTable(&models.Listaprecioh{})
	db.Migrator().DropTable(&models.TipoMovinv{})
	db.Migrator().DropTable(&models.FormaPago{})
	db.Migrator().DropTable(&models.Producto{})
	db.Migrator().DropTable(&models.Proveedor{})
	db.Migrator().DropTable(&models.ClienteDireccion{})
	db.Migrator().DropTable(&models.ClienteContacto{})
	db.Migrator().DropTable(&models.TipoContacto{})
	db.Migrator().DropTable(&models.Municipio{})
	db.Migrator().DropTable(&models.Departamento{})
	db.Migrator().DropTable(&models.Cliente{})
	db.Migrator().DropTable(&models.Empleado{})
	db.Migrator().DropTable(&models.Sucursal{})
	db.Migrator().DropTable(&models.Empresa{})
	db.Migrator().DropTable(&models.Pais{})
	db.Migrator().DropTable(&models.Moneda{})
	db.Migrator().DropTable(&models.Sysparam{})

	db.AutoMigrate(&models.Sysparam{})
	db.AutoMigrate(&models.Moneda{})
	db.AutoMigrate(&models.Pais{})
	db.AutoMigrate(&models.Empresa{})
	db.AutoMigrate(&models.Departamento{})
	db.AutoMigrate(&models.Municipio{})
	db.AutoMigrate(&models.Sucursal{})
	db.AutoMigrate(&models.Empleado{})
	db.AutoMigrate(&models.Cliente{})
	db.AutoMigrate(&models.ClienteDireccion{})
	db.AutoMigrate(&models.TipoContacto{})
	db.AutoMigrate(&models.ClienteContacto{})
	db.AutoMigrate(&models.Proveedor{})
	db.AutoMigrate(&models.Producto{})
	db.AutoMigrate(&models.FormaPago{})
	db.AutoMigrate(&models.TipoMovinv{})
	db.AutoMigrate(&models.Listaprecioh{})
	db.AutoMigrate(&models.Listapreciod{})
	db.AutoMigrate(&models.Movinvh{})
	db.AutoMigrate(&models.Movinvd{})
	db.AutoMigrate(&models.ClienteContador{})
	db.AutoMigrate(&models.ClienteLectura{})
	db.AutoMigrate(&models.Caja{})
	db.AutoMigrate(&models.CajaMov{})
	db.AutoMigrate(&models.Perfil{})
	db.AutoMigrate(&models.Usuario{})
	db.AutoMigrate(&models.Ordenh{})
	db.AutoMigrate(&models.Ordend{})
	db.AutoMigrate(&models.Menu{})

	return db
}
