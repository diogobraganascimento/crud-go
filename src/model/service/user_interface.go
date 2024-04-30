package service

import (
	"github.com/diogobraganascimento/crud-go/src/configuration/rest_err"
	"github.com/diogobraganascimento/crud-go/src/model"
	"github.com/diogobraganascimento/crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
