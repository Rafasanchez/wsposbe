package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
type handler struct {
	DB *gorm.DB
}*/

func (h handler) findMoneda(c *gin.Context) {
	id := c.Param("id")

	var Moneda models.Moneda

	if result := h.DB.First(&Moneda, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Moneda})
}

func (h handler) findMonedas(c *gin.Context) {
	var Moneda []models.Moneda

	if result := h.DB.Find(&Moneda); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Moneda})
}

func (h handler) createMoneda(c *gin.Context) {
	var body models.Moneda

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Moneda models.Moneda

	Moneda.Moneda = body.Moneda
	Moneda.CodigoMoneda = body.CodigoMoneda
	Moneda.Estado = body.Estado

	if result := h.DB.Create(&Moneda); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateMoneda(c *gin.Context) {
	id := c.Param("id")
	var body models.Moneda

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Moneda models.Moneda

	if result := h.DB.First(&Moneda, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Moneda.Moneda = body.Moneda
	Moneda.CodigoMoneda = body.CodigoMoneda
	Moneda.Estado = body.Estado

	h.DB.Save(&Moneda)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteMoneda(c *gin.Context) {
	id := c.Param("id")

	var Moneda models.Moneda

	if result := h.DB.First(&Moneda, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Moneda)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesMoneda(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/moneda")
	routes.POST("/", h.createMoneda)
	routes.GET("/", h.findMonedas)
	routes.GET("/:id", h.findMoneda)
	routes.PUT("/:id", h.updateMoneda)
	routes.DELETE("/:id", h.deleteMoneda)
}
