package controller

import (
	"net/http"

	"github.com/diogobraganascimento/crud-go/src/configuration/logger"
	"github.com/diogobraganascimento/crud-go/src/configuration/validation"
	"github.com/diogobraganascimento/crud-go/src/controller/model/request"
	"github.com/diogobraganascimento/crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zapcore.Field{
				Key:    "journey",
				String: "createUser",
			})
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info(
		"User created successfully",
		zapcore.Field{
			Key:    "journey",
			String: "CreateUser",
		})

	c.JSON(http.StatusOK, response)
}
