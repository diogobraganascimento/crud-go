package service

import (
	"github.com/diogobraganascimento/crud-go/src/configuration/logger"
	"github.com/diogobraganascimento/crud-go/src/configuration/rest_err"
	"github.com/diogobraganascimento/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init loginUser model.",
		zap.String("journey", "loginUser"))

	user, err := ud.findUserByEmailAndPasswordServices(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, err
	}

	userDomain.EncryptPassword()

	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"))
	return user, nil

}
