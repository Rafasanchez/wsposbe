package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findSysparam(c *gin.Context) {
	id := c.Param("id")

	var Sysparam models.Sysparam

	if result := h.DB.First(&Sysparam, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Sysparam})
}

func (h handler) findSysparams(c *gin.Context) {
	var Sysparam []models.Sysparam

	if result := h.DB.Find(&Sysparam); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Sysparam})
}

func (h handler) createSysparam(c *gin.Context) {
	var body models.Sysparam

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Sysparam models.Sysparam

	Sysparam.Parametro = body.Parametro
	Sysparam.Descripcion = body.Descripcion
	Sysparam.Valor = body.Valor
	Sysparam.Tipo = body.Tipo
	Sysparam.Estado = body.Estado

	if result := h.DB.Create(&Sysparam); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1002", "success": false, "message": functions.GetInfoMsg(1002), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) updateSysparam(c *gin.Context) {
	id := c.Param("id")
	var body models.Sysparam

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetInfoMsg(9999), "errors": functions.GetErrorMsg(err)})
		return
	}

	var Sysparam models.Sysparam

	if result := h.DB.First(&Sysparam, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	Sysparam.Parametro = body.Parametro
	Sysparam.Descripcion = body.Descripcion
	Sysparam.Valor = body.Valor
	Sysparam.Tipo = body.Tipo
	Sysparam.Estado = body.Estado

	h.DB.Save(&Sysparam)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) deleteSysparam(c *gin.Context) {
	id := c.Param("id")

	var Sysparam models.Sysparam

	if result := h.DB.First(&Sysparam, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	h.DB.Delete(&Sysparam)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000)})
}

func (h handler) getParam(c *gin.Context) {
	paramcode := c.Param("paramcode")

	var Sysparam models.Sysparam

	if result := h.DB.Select("valor").Find(&Sysparam, "parametro = ?", paramcode); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": functions.GetInfoMsg(1001), "errors": functions.GetErrorMsg(result.Error)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": functions.GetInfoMsg(1000), "response": &Sysparam})
}

func RegisterRoutesSysparam(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/sysparam")
	routes.POST("/", h.createSysparam)
	routes.GET("/", h.findSysparams)
	routes.GET("/:id", h.findSysparam)
	routes.PUT("/:id", h.updateSysparam)
	routes.DELETE("/:id", h.deleteSysparam)
	routes.GET("/getParam/:paramcode", h.getParam)
}
