package controllers

import (
	"net/http"
	"wsposbe/pkg/common/functions"
	"wsposbe/pkg/common/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) findCliente(c *gin.Context) {
	id := c.Param("id")

	var Cliente models.Cliente

	if result := h.DB.First(&Cliente, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, &Cliente)
}

func (h handler) findClientes(c *gin.Context) {
	var Cliente []models.Cliente

	if result := h.DB.Find(&Cliente); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}

	c.JSON(http.StatusOK, &Cliente)
}

func (h handler) createCliente(c *gin.Context) {
	var body models.Cliente

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": functions.GetErrorMsg(err)})
		return
	}

	var Cliente models.Cliente

	Cliente.IDEmpresa = body.IDEmpresa
	Cliente.PrimerNombre = body.PrimerNombre
	Cliente.SegundoNombre = body.SegundoNombre
	Cliente.PrimerApellido = body.PrimerApellido
	Cliente.SegundoApellido = body.SegundoApellido
	Cliente.ApellidoCasada = body.ApellidoCasada
	Cliente.CorreoElectronico = body.CorreoElectronico
	Cliente.FechaNacimiento = body.FechaNacimiento
	Cliente.Nit = body.Nit
	Cliente.DPI = body.DPI
	Cliente.Telefono = body.Telefono
	Cliente.Estado = body.Estado

	if result := h.DB.Create(&Cliente); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}

	c.JSON(http.StatusCreated, &Cliente)
}

func (h handler) updateCliente(c *gin.Context) {
	id := c.Param("id")
	var body models.Cliente

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": functions.GetErrorMsg(err)})
		return
	}

	var Cliente models.Cliente

	if result := h.DB.First(&Cliente, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}

	Cliente.IDEmpresa = body.IDEmpresa
	Cliente.PrimerNombre = body.PrimerNombre
	Cliente.SegundoNombre = body.SegundoNombre
	Cliente.PrimerApellido = body.PrimerApellido
	Cliente.SegundoApellido = body.SegundoApellido
	Cliente.ApellidoCasada = body.ApellidoCasada
	Cliente.CorreoElectronico = body.CorreoElectronico
	Cliente.FechaNacimiento = body.FechaNacimiento
	Cliente.Nit = body.Nit
	Cliente.DPI = body.DPI
	Cliente.Telefono = body.Telefono
	Cliente.Estado = body.Estado

	h.DB.Save(&Cliente)

	c.JSON(http.StatusOK, &Cliente)
}

func (h handler) deleteCliente(c *gin.Context) {
	id := c.Param("id")

	var Cliente models.Cliente

	if result := h.DB.First(&Cliente, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}

	h.DB.Delete(&Cliente)

	c.Status(http.StatusOK)
}

func RegisterRoutesCliente(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/cliente")
	routes.POST("/", h.createCliente)
	routes.GET("/", h.findClientes)
	routes.GET("/:id", h.findCliente)
	routes.PUT("/:id", h.updateCliente)
	routes.DELETE("/:id", h.deleteCliente)
}
