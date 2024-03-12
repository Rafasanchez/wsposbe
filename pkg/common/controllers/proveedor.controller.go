package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findProveedor(c *gin.Context) {
	id := c.Param("id")

	var Proveedor models.Proveedor

	if result := h.DB.First(&Proveedor, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Proveedor})
}

func (h handler) findProveedores(c *gin.Context) {
	var Proveedor []models.Proveedor

	if result := h.DB.Find(&Proveedor); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Proveedor})
}

func (h handler) createProveedor(c *gin.Context) {
	var body models.Proveedor

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Proveedor models.Proveedor

	Proveedor.IDEmpresa = body.IDEmpresa
	Proveedor.Nombres = body.Nombres
	Proveedor.Apellidos = body.Apellidos
	Proveedor.CorreoElectronico = body.CorreoElectronico
	Proveedor.Nit = body.Nit
	Proveedor.DPI = body.DPI
	Proveedor.Direccion = body.Direccion
	Proveedor.Telefono = body.Telefono
	Proveedor.Estado = body.Estado

	if result := h.DB.Create(&Proveedor); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateProveedor(c *gin.Context) {
	id := c.Param("id")
	var body models.Proveedor

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Proveedor models.Proveedor

	if result := h.DB.First(&Proveedor, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Proveedor.IDEmpresa = body.IDEmpresa
	Proveedor.Nombres = body.Nombres
	Proveedor.Apellidos = body.Apellidos
	Proveedor.CorreoElectronico = body.CorreoElectronico
	Proveedor.Nit = body.Nit
	Proveedor.DPI = body.DPI
	Proveedor.Direccion = body.Direccion
	Proveedor.Telefono = body.Telefono
	Proveedor.Estado = body.Estado

	h.DB.Save(&Proveedor)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteProveedor(c *gin.Context) {
	id := c.Param("id")

	var Proveedor models.Proveedor

	if result := h.DB.First(&Proveedor, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Proveedor)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesProveedor(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/proveedor")
	routes.POST("/", h.createProveedor)
	routes.GET("/", h.findProveedores)
	routes.GET("/:id", h.findProveedor)
	routes.PUT("/:id", h.updateProveedor)
	routes.DELETE("/:id", h.deleteProveedor)
}
