package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findPais(c *gin.Context) {
	id := c.Param("id")

	var Pais models.Pais

	if result := h.DB.First(&Pais, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Pais})
}

func (h handler) findPaises(c *gin.Context) {
	var Pais []models.Pais

	if result := h.DB.Find(&Pais); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Pais})
}

func (h handler) createPais(c *gin.Context) {
	var body models.Pais

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Pais models.Pais

	Pais.Pais = body.Pais
	Pais.CodigoPais = body.CodigoPais
	Pais.IDMoneda = body.IDMoneda
	Pais.Estado = body.Estado

	if result := h.DB.Create(&Pais); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updatePais(c *gin.Context) {
	id := c.Param("id")
	var body models.Pais

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Pais models.Pais

	if result := h.DB.First(&Pais, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Pais.Pais = body.Pais
	Pais.CodigoPais = body.CodigoPais
	Pais.IDMoneda = body.IDMoneda
	Pais.Estado = body.Estado

	h.DB.Save(&Pais)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deletePais(c *gin.Context) {
	id := c.Param("id")

	var Pais models.Pais

	if result := h.DB.First(&Pais, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Pais)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesPais(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/pais")
	routes.POST("/", h.createPais)
	routes.GET("/", h.findPaises)
	routes.GET("/:id", h.findPais)
	routes.PUT("/:id", h.updatePais)
	routes.DELETE("/:id", h.deletePais)
}
