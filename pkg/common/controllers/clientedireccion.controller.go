package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findClienteDireccion(c *gin.Context) {
	id := c.Param("id")

	var ClienteDireccion models.ClienteDireccion

	if result := h.DB.First(&ClienteDireccion, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "", "Response": &ClienteDireccion})
}

func (h handler) findClienteDirecciones(c *gin.Context) {
	var ClienteDireccion []models.ClienteDireccion

	if result := h.DB.Find(&ClienteDireccion); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "", "Response": &ClienteDireccion})
}

func (h handler) createClienteDireccion(c *gin.Context) {
	var body models.ClienteDireccion

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var ClienteDireccion models.ClienteDireccion

	ClienteDireccion.IDMunicipio = body.IDMunicipio
	ClienteDireccion.IDCliente = body.IDCliente
	ClienteDireccion.Direccion = body.Direccion
	ClienteDireccion.Latitud = body.Latitud
	ClienteDireccion.Longitud = body.Longitud
	ClienteDireccion.Estado = body.Estado

	if result := h.DB.Create(&ClienteDireccion); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": "No se pudo crear la dirección", "errors": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": "Tipo Contacto creado"})
}

func (h handler) updateClienteDireccion(c *gin.Context) {
	id := c.Param("id")
	var body models.ClienteDireccion

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var ClienteDireccion models.ClienteDireccion

	if result := h.DB.First(&ClienteDireccion, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": "No se encontro la dirección", "errors": result.Error})
		return
	}

	ClienteDireccion.IDMunicipio = body.IDMunicipio
	ClienteDireccion.IDCliente = body.IDCliente
	ClienteDireccion.Direccion = body.Direccion
	ClienteDireccion.Latitud = body.Latitud
	ClienteDireccion.Longitud = body.Longitud
	ClienteDireccion.Estado = body.Estado

	h.DB.Save(&ClienteDireccion)

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": "Tipo Contacto actualizado"})
}

func (h handler) deleteClienteDireccion(c *gin.Context) {
	id := c.Param("id")

	var ClienteDireccion models.ClienteDireccion

	if result := h.DB.First(&ClienteDireccion, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": "No se encontro el producto", "errors": result.Error})
		return
	}

	h.DB.Delete(&ClienteDireccion)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "Cliente Dirección eliminado"})
}

func RegisterRoutesClienteDireccion(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/clientedireccion")
	routes.POST("/", h.createClienteDireccion)
	routes.GET("/", h.findClienteDirecciones)
	routes.GET("/:id", h.findClienteDireccion)
	routes.PUT("/:id", h.updateClienteDireccion)
	routes.DELETE("/:id", h.deleteClienteDireccion)
}
