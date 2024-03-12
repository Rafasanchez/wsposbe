package main

import (
	"wsposbe/pkg/common/controllers"
	"wsposbe/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	controllers.RegisterRoutesSysparam(r, h)
	controllers.RegisterRoutesMoneda(r, h)
	controllers.RegisterRoutesPais(r, h)
	controllers.RegisterRoutesEmpresa(r, h)
	controllers.RegisterRoutesDepartamento(r, h)
	controllers.RegisterRoutesMunicipio(r, h)
	controllers.RegisterRoutesTipoContacto(r, h)
	controllers.RegisterRoutesCliente(r, h)
	controllers.RegisterRoutesCaja(r, h)
	controllers.RegisterRoutesFormaPago(r, h)
	controllers.RegisterRoutesSucursal(r, h)
	controllers.RegisterRoutesTipoMovinv(r, h)
	controllers.RegisterRoutesProveedor(r, h)
	controllers.RegisterRoutesCajaMov(r, h)
	controllers.RegisterRoutesClienteDireccion(r, h)
	controllers.RegisterRoutesProducto(r, h)

	r.Run(port)
}
