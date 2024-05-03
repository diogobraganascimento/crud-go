package converter

import (
	"github.com/diogobraganascimento/crud-go/src/model"
	"github.com/diogobraganascimento/crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
