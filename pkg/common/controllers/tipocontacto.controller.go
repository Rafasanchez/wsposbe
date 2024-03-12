package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findTipoContacto(c *gin.Context) {
	id := c.Param("id")

	var TipoContacto models.TipoContacto

	if result := h.DB.First(&TipoContacto, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &TipoContacto})
}

func (h handler) findTiposContacto(c *gin.Context) {
	var TipoContacto []models.TipoContacto

	if result := h.DB.Find(&TipoContacto); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &TipoContacto})
}

func (h handler) createTipoContacto(c *gin.Context) {
	var body models.TipoContacto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var TipoContacto models.TipoContacto

	TipoContacto.Codigo = body.Codigo
	TipoContacto.Descripcion = body.Descripcion
	TipoContacto.Estado = body.Estado

	if result := h.DB.Create(&TipoContacto); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateTipoContacto(c *gin.Context) {
	id := c.Param("id")
	var body models.TipoContacto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var TipoContacto models.TipoContacto

	if result := h.DB.First(&TipoContacto, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	TipoContacto.Codigo = body.Codigo
	TipoContacto.Descripcion = body.Descripcion
	TipoContacto.Estado = body.Estado

	h.DB.Save(&TipoContacto)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteTipoContacto(c *gin.Context) {
	id := c.Param("id")

	var TipoContacto models.TipoContacto

	if result := h.DB.First(&TipoContacto, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&TipoContacto)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesTipoContacto(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/tipocontacto")
	routes.POST("/", h.createTipoContacto)
	routes.GET("/", h.findTiposContacto)
	routes.GET("/:id", h.findTipoContacto)
	routes.PUT("/:id", h.updateTipoContacto)
	routes.DELETE("/:id", h.deleteTipoContacto)
}
