package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findProducto(c *gin.Context) {
	id := c.Param("id")

	var Producto models.Producto

	if result := h.DB.First(&Producto, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Producto})
}

func (h handler) findProductos(c *gin.Context) {
	var Producto []models.Producto

	if result := h.DB.Find(&Producto); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Producto})
}

func (h handler) createProducto(c *gin.Context) {
	var body models.Producto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Producto models.Producto

	Producto.IDEmpresa = body.IDEmpresa
	Producto.Codigo = body.Codigo
	Producto.Descripcion = body.Descripcion
	Producto.Imagen = body.Imagen
	Producto.CobroxTiempo = body.CobroxTiempo
	Producto.GeneraTicket = body.GeneraTicket
	Producto.CargoParqueo = body.CargoParqueo
	Producto.Estado = body.Estado

	if result := h.DB.Create(&Producto); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateProducto(c *gin.Context) {
	id := c.Param("id")
	var body models.Producto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Producto models.Producto

	if result := h.DB.First(&Producto, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Producto.IDEmpresa = body.IDEmpresa
	Producto.Codigo = body.Codigo
	Producto.Descripcion = body.Descripcion
	Producto.Imagen = body.Imagen
	Producto.CobroxTiempo = body.CobroxTiempo
	Producto.GeneraTicket = body.GeneraTicket
	Producto.CargoParqueo = body.CargoParqueo
	Producto.Estado = body.Estado

	h.DB.Save(&Producto)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteProducto(c *gin.Context) {
	id := c.Param("id")

	var Producto models.Producto

	if result := h.DB.First(&Producto, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Producto)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesProducto(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/producto")
	routes.POST("/", h.createProducto)
	routes.GET("/", h.findProductos)
	routes.GET("/:id", h.findProducto)
	routes.PUT("/:id", h.updateProducto)
	routes.DELETE("/:id", h.deleteProducto)
}
