package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findDepartamento(c *gin.Context) {
	id := c.Param("id")

	var Departamento models.Departamento

	if result := h.DB.First(&Departamento, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Departamento})
}

func (h handler) findDepartamentos(c *gin.Context) {
	var Departamento []models.Departamento

	if result := h.DB.Find(&Departamento); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Departamento})
}

func (h handler) createDepartamento(c *gin.Context) {
	var body models.Departamento

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Departamento models.Departamento

	Departamento.Departamento = body.Departamento
	Departamento.IDPais = body.IDPais
	Departamento.Estado = body.Estado

	if result := h.DB.Create(&Departamento); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateDepartamento(c *gin.Context) {
	id := c.Param("id")
	var body models.Departamento

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Departamento models.Departamento

	if result := h.DB.First(&Departamento, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Departamento.Departamento = body.Departamento
	Departamento.IDPais = body.IDPais
	Departamento.Estado = body.Estado

	h.DB.Save(&Departamento)

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteDepartamento(c *gin.Context) {
	id := c.Param("id")

	var Departamento models.Departamento

	if result := h.DB.First(&Departamento, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Departamento)

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesDepartamento(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/departamento")
	routes.POST("/", h.createDepartamento)
	routes.GET("/", h.findDepartamentos)
	routes.GET("/:id", h.findDepartamento)
	routes.PUT("/:id", h.updateDepartamento)
	routes.DELETE("/:id", h.deleteDepartamento)
}
