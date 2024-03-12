package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findEmpresa(c *gin.Context) {
	id := c.Param("id")

	var Empresa models.Empresa

	if result := h.DB.First(&Empresa, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Empresa})
}

func (h handler) findEmpresas(c *gin.Context) {
	var Empresa []models.Empresa

	if result := h.DB.Find(&Empresa); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Empresa})
}

func (h handler) createEmpresa(c *gin.Context) {
	var body models.Empresa

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Empresa models.Empresa

	Empresa.Empresa = body.Empresa
	Empresa.Direccion = body.Direccion
	Empresa.Nit = body.Nit
	Empresa.Telefono = body.Telefono
	Empresa.IDPais = body.IDPais
	Empresa.Estado = body.Estado

	if result := h.DB.Create(&Empresa); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateEmpresa(c *gin.Context) {
	id := c.Param("id")
	var body models.Empresa

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Empresa models.Empresa

	if result := h.DB.First(&Empresa, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Empresa.Empresa = body.Empresa
	Empresa.Direccion = body.Direccion
	Empresa.Nit = body.Nit
	Empresa.Telefono = body.Telefono
	Empresa.IDPais = body.IDPais
	Empresa.Estado = body.Estado

	h.DB.Save(&Empresa)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteEmpresa(c *gin.Context) {
	id := c.Param("id")

	var Empresa models.Empresa

	if result := h.DB.First(&Empresa, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Empresa)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesEmpresa(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/empresa")
	routes.POST("/", h.createEmpresa)
	routes.GET("/", h.findEmpresas)
	routes.GET("/:id", h.findEmpresa)
	routes.PUT("/:id", h.updateEmpresa)
	routes.DELETE("/:id", h.deleteEmpresa)
}
