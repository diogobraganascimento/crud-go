package service

import (
	"github.com/diogobraganascimento/crud-go/src/configuration/rest_err"
	"github.com/diogobraganascimento/crud-go/src/model"
)

func (*userDomainService) UpdateUser(
	userID string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
