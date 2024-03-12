package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findMunicipio(c *gin.Context) {
	id := c.Param("id")

	var Municipio models.Municipio

	if result := h.DB.First(&Municipio, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Municipio})
}

func (h handler) findMunicipios(c *gin.Context) {
	var Municipio []models.Municipio

	if result := h.DB.Find(&Municipio); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Municipio})
}

func (h handler) createMunicipio(c *gin.Context) {
	var body models.Municipio

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Municipio models.Municipio

	Municipio.IDDepartamento = body.IDDepartamento
	Municipio.Municipio = body.Municipio
	Municipio.Estado = body.Estado

	if result := h.DB.Create(&Municipio); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateMunicipio(c *gin.Context) {
	id := c.Param("id")
	var body models.Municipio

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Municipio models.Municipio

	if result := h.DB.First(&Municipio, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Municipio.IDDepartamento = body.IDDepartamento
	Municipio.Municipio = body.Municipio
	Municipio.Estado = body.Estado

	h.DB.Save(&Municipio)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteMunicipio(c *gin.Context) {
	id := c.Param("id")

	var Municipio models.Municipio

	if result := h.DB.First(&Municipio, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Municipio)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesMunicipio(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/municipio")
	routes.POST("/", h.createMunicipio)
	routes.GET("/", h.findMunicipios)
	routes.GET("/:id", h.findMunicipio)
	routes.PUT("/:id", h.updateMunicipio)
	routes.DELETE("/:id", h.deleteMunicipio)
}
