package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findFormaPago(c *gin.Context) {
	id := c.Param("id")

	var FormaPago models.FormaPago

	if result := h.DB.First(&FormaPago, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &FormaPago})
}

func (h handler) findFormasPagos(c *gin.Context) {
	var FormaPago []models.FormaPago

	if result := h.DB.Find(&FormaPago); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &FormaPago})
}

func (h handler) createFormaPago(c *gin.Context) {
	var body models.FormaPago

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var FormaPago models.FormaPago

	FormaPago.IDEmpresa = body.IDEmpresa
	FormaPago.Codigo = body.Codigo
	FormaPago.Descripcion = body.Descripcion
	FormaPago.Estado = body.Estado

	if result := h.DB.Create(&FormaPago); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateFormaPago(c *gin.Context) {
	id := c.Param("id")
	var body models.FormaPago

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var FormaPago models.FormaPago

	if result := h.DB.First(&FormaPago, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	FormaPago.IDEmpresa = body.IDEmpresa
	FormaPago.Codigo = body.Codigo
	FormaPago.Descripcion = body.Descripcion
	FormaPago.Estado = body.Estado

	h.DB.Save(&FormaPago)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteFormaPago(c *gin.Context) {
	id := c.Param("id")

	var FormaPago models.FormaPago

	if result := h.DB.First(&FormaPago, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&FormaPago)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func RegisterRoutesFormaPago(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/formapago")
	routes.POST("/", h.createFormaPago)
	routes.GET("/", h.findFormasPagos)
	routes.GET("/:id", h.findFormaPago)
	routes.PUT("/:id", h.updateFormaPago)
	routes.DELETE("/:id", h.deleteFormaPago)
}
