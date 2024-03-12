package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findSucursal(c *gin.Context) {
	id := c.Param("id")

	var Sucursal models.Sucursal

	if result := h.DB.First(&Sucursal, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Sucursal})
}

func (h handler) findSucursales(c *gin.Context) {
	var Sucursal []models.Sucursal

	if result := h.DB.Find(&Sucursal); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Sucursal})
}

func (h handler) createSucursal(c *gin.Context) {
	var body models.Sucursal

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Sucursal models.Sucursal

	Sucursal.IDEmpresa = body.IDEmpresa
	Sucursal.Codigo = body.Codigo
	Sucursal.Descripcion = body.Descripcion
	Sucursal.Direccion = body.Direccion
	Sucursal.Telefono = body.Telefono
	Sucursal.Estado = body.Estado

	if result := h.DB.Create(&Sucursal); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateSucursal(c *gin.Context) {
	id := c.Param("id")
	var body models.Sucursal

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Sucursal models.Sucursal

	if result := h.DB.First(&Sucursal, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Sucursal.IDEmpresa = body.IDEmpresa
	Sucursal.Codigo = body.Codigo
	Sucursal.Descripcion = body.Descripcion
	Sucursal.Direccion = body.Direccion
	Sucursal.Telefono = body.Telefono
	Sucursal.Estado = body.Estado

	h.DB.Save(&Sucursal)

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteSucursal(c *gin.Context) {
	id := c.Param("id")

	var Sucursal models.Sucursal

	if result := h.DB.First(&Sucursal, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Sucursal)

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesSucursal(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/sucursal")
	routes.POST("/", h.createSucursal)
	routes.GET("/", h.findSucursales)
	routes.GET("/:id", h.findSucursal)
	routes.PUT("/:id", h.updateSucursal)
	routes.DELETE("/:id", h.deleteSucursal)
}
