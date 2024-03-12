package controllers

import (
	"fmt"
	"net/http"
	"time"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type listaprecio struct {
	IDListaprecioh uint
	IDEmpresa      uint
	Codigo         string
	Descripcion    string
	FechaDesde     time.Time
	FechaHasta     time.Time
	Estado         uint8
	IDProducto     uint
	Precio         float32
}

func (h handler) findListaPrecio(c *gin.Context) {
	id := c.Param("id")

	var Listaprecioh models.Listaprecioh

	if result := h.DB.InnerJoins("Listapreciod").First(&Listaprecioh, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Listaprecioh})
}

func (h handler) findListasPrecio(c *gin.Context) {
	var Listaprecioh []models.Listaprecioh

	if result := h.DB.Find(&Listaprecioh); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Listaprecioh})
}

func (h handler) createListaPrecio(c *gin.Context) {
	fmt.Println(c.Request.Body)

	var body listaprecio

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}
	/*
		var Listaprecioh models.Listaprecioh

		if result := h.DB.First(&Listaprecioh, body.IDListaprecioh); result.Error != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": result.Error})
			return
		}

		Listaprecioh.ID = body.IDListaprecioh
		Listaprecioh.IDEmpresa = body.IDEmpresa
		Listaprecioh.Codigo = body.Codigo
		Listaprecioh.Descripcion = body.Descripcion
		Listaprecioh.FechaDesde = body.FechaDesde
		Listaprecioh.FechaHasta = body.FechaHasta
		Listaprecioh.Estado = body.Estado

		if result := h.DB.Create(&Listaprecioh); result.Error != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
			return
		}

		var Listapreciod models.Listapreciod

		values := reflect.ValueOf(body)
		types := values.Type()
		for i := 0; i < values.NumField(); i++ {
			Listapreciod.IDListaprecioh = types.Field(i).IDListaprecioh
			Listapreciod.IDProducto = types.Field(i).IDProducto
			Listapreciod.Precio = types.Field(i).Precio

			if result := h.DB.Create(&Listapreciod); result.Error != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
				return
			}
		}

		c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
	*/
}

func RegisterRoutesListaPrecio(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/listaprecio")
	routes.POST("/", h.createListaPrecio)
	routes.GET("/", h.findListasPrecio)
	routes.GET("/:id", h.findListaPrecio)
	//routes.PUT("/:id", h.updateListaPrecio)
	//routes.DELETE("/:id", h.deleteListaPrecio)
}
