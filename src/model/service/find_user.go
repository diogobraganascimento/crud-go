package service

import (
	"github.com/diogobraganascimento/crud-go/src/configuration/rest_err"
	"github.com/diogobraganascimento/crud-go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
