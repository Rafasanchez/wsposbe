package controllers

import (
	"net/http"
	"time"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findCajaMov(c *gin.Context) {
	idcaja := c.Param("idcaja")

	var Caja models.Caja
	var CajaMov models.CajaMov
	now := time.Now()

	if result := h.DB.First(&Caja, idcaja); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "1001", "success": false, "message": "La caja no existe"})
		return
	}

	result := h.DB.Where("id_caja = ? and fecha = ?", idcaja, now.Round(0)).Find(&CajaMov)

	println(string(result.Statement.Table))

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "1002", "success": false, "message": "La caja no esta abierta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "", "Response": &CajaMov})
}

func (h handler) aperturaCaja(c *gin.Context) {
	var body models.CajaMov
	var CajaMov models.CajaMov

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetErrorMsg(err)})
		return
	}

	result := h.DB.Where("id_caja = ? and fecha = ?", body.IDCaja, functions.Today()).Find(&CajaMov)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	if result.RowsAffected > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "1003", "success": false, "message": "La caja ya fue abierta"})
		return
	}

	CajaMov.IDCaja = body.IDCaja
	CajaMov.Fecha = functions.Today()
	CajaMov.HoraApertura = time.Now()
	CajaMov.MontoApertura = body.MontoApertura
	CajaMov.Estado = 1

	if result := h.DB.Create(&CajaMov); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": "1000", "success": true, "message": "Caja aperturada"})
}

func (h handler) cierreCaja(c *gin.Context) {
	var body models.CajaMov
	var CajaMov models.CajaMov

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "9999", "success": false, "message": functions.GetErrorMsg(err)})
		return
	}

	result := h.DB.Where("id_caja = ? and fecha = ?", body.IDCaja, functions.Today()).Find(&CajaMov)

	println(result.PrepareStmt)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "9999", "success": false, "message": result.Error})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "1004", "success": false, "message": "La caja no esta abierta"})
		return
	}

	CajaMov.HoraCierre = time.Now()
	CajaMov.MontoCierre = body.MontoCierre
	CajaMov.Estado = 0

	h.DB.Save(&CajaMov)

	c.JSON(http.StatusOK, gin.H{"code": "1000", "success": true, "message": "Caja cerrada"})
}

func RegisterRoutesCajaMov(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/cajamov")
	routes.POST("/", h.aperturaCaja)
	routes.GET("/:idcaja", h.findCajaMov)
	routes.PUT("/", h.cierreCaja)
}
