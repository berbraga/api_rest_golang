package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"numismatico/api/entities"
)

type cedulaController struct {
	cedulas []entities.Cedula
}

func NewCedulaController() *cedulaController {
	return &cedulaController{}
}

func (c *cedulaController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.cedulas)
}
func (c *cedulaController) Create(ctx *gin.Context) {
	cedula := entities.NewCedula()
	if err := ctx.BindJSON(&cedula); err != nil {
		return
	}

	c.cedulas = append(c.cedulas, *cedula)

	ctx.JSON(http.StatusOK, c.cedulas)
}
func (c *cedulaController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	for idx, cedula := range c.cedulas {
		if cedula.ID == id {
			c.cedulas = append(c.cedulas[0:idx], c.cedulas[idx+1:]...)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "cedula nao encontrada"})
}
