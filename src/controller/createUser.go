package controller

import (
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Voce chamaou a rota de forma errada")
	c.JSON(err.Code, err)
}
