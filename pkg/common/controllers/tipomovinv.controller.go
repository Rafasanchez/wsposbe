package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findTipoMovinv(c *gin.Context) {
	id := c.Param("id")

	var TipoMovinv models.TipoMovinv

	if result := h.DB.First(&TipoMovinv, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &TipoMovinv})
}

func (h handler) findTiposMovinv(c *gin.Context) {
	var TipoMovinv []models.TipoMovinv

	if result := h.DB.Find(&TipoMovinv); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &TipoMovinv})
}

func (h handler) createTipoMovinv(c *gin.Context) {
	var body models.TipoMovinv

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var TipoMovinv models.TipoMovinv

	TipoMovinv.IDEmpresa = body.IDEmpresa
	TipoMovinv.Codigo = body.Codigo
	TipoMovinv.Descripcion = body.Descripcion
	TipoMovinv.SumaInventario = body.SumaInventario
	TipoMovinv.SumaCosto = body.SumaCosto
	TipoMovinv.Estado = body.Estado

	if result := h.DB.Create(&TipoMovinv); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateTipoMovinv(c *gin.Context) {
	id := c.Param("id")
	var body models.TipoMovinv

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var TipoMovinv models.TipoMovinv

	if result := h.DB.First(&TipoMovinv, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	TipoMovinv.IDEmpresa = body.IDEmpresa
	TipoMovinv.Codigo = body.Codigo
	TipoMovinv.Descripcion = body.Descripcion
	TipoMovinv.SumaInventario = body.SumaInventario
	TipoMovinv.SumaCosto = body.SumaCosto
	TipoMovinv.Estado = body.Estado

	h.DB.Save(&TipoMovinv)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteTipoMovinv(c *gin.Context) {
	id := c.Param("id")

	var TipoMovinv models.TipoMovinv

	if result := h.DB.First(&TipoMovinv, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&TipoMovinv)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesTipoMovinv(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/tipomovinv")
	routes.POST("/", h.createTipoMovinv)
	routes.GET("/", h.findTiposMovinv)
	routes.GET("/:id", h.findTipoMovinv)
	routes.PUT("/:id", h.updateTipoMovinv)
	routes.DELETE("/:id", h.deleteTipoMovinv)
}
