package controller

import (
	"net/http"
	"strings"

	"github.com/diogobraganascimento/crud-go/src/configuration/logger"
	"github.com/diogobraganascimento/crud-go/src/configuration/validation"
	"github.com/diogobraganascimento/crud-go/src/controller/model/request"
	"github.com/diogobraganascimento/crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest

	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)
	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error(
			"error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"updateUser controller executed sucessfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
