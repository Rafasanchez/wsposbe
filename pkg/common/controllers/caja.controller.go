package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findCaja(c *gin.Context) {
	id := c.Param("id")

	var Caja models.Caja

	if result := h.DB.First(&Caja, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	//c.JSON(http.StatusOK, &Caja)
	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "", "Response": &Caja})
}

func (h handler) findCajas(c *gin.Context) {
	var Caja []models.Caja

	if result := h.DB.Find(&Caja); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "", "Response": &Caja})
}

func (h handler) createCaja(c *gin.Context) {
	var body models.Caja

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetErrorMsg(err)})
		return
	}

	var Caja models.Caja

	Caja.IDEmpresa = body.IDEmpresa
	Caja.TipoAutorizacion = body.TipoAutorizacion
	Caja.FechaAutorizacion = body.FechaAutorizacion
	Caja.Resolucion = body.Resolucion
	Caja.Serie = body.Serie
	Caja.InicioCorrelativo = body.InicioCorrelativo
	Caja.FinalCorrelativo = body.FinalCorrelativo
	Caja.Estado = body.Estado

	if result := h.DB.Create(&Caja); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": "Caja creada"})
}

func (h handler) updateCaja(c *gin.Context) {
	id := c.Param("id")
	var body models.Caja

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetErrorMsg(err)})
		return
	}

	var Caja models.Caja

	if result := h.DB.First(&Caja, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	Caja.IDEmpresa = body.IDEmpresa
	Caja.TipoAutorizacion = body.TipoAutorizacion
	Caja.FechaAutorizacion = body.FechaAutorizacion
	Caja.Resolucion = body.Resolucion
	Caja.Serie = body.Serie
	Caja.InicioCorrelativo = body.InicioCorrelativo
	Caja.FinalCorrelativo = body.FinalCorrelativo
	Caja.Estado = body.Estado

	h.DB.Save(&Caja)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "Caja modificada"})
}

func (h handler) deleteCaja(c *gin.Context) {
	id := c.Param("id")

	var Caja models.Caja

	if result := h.DB.First(&Caja, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	h.DB.Delete(&Caja)

	//c.Status(http.StatusOK)
	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "Caja eliminada"})
}

func RegisterRoutesCaja(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/caja")
	routes.POST("/", h.createCaja)
	routes.GET("/", h.findCajas)
	routes.GET("/:id", h.findCaja)
	routes.PUT("/:id", h.updateCaja)
	routes.DELETE("/:id", h.deleteCaja)
}
