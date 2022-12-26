package userrepository

import (
	"service-campaign-startup/model/entity"
)

type UserRepository interface {
	RegisterUser(entity.User) (entity.User, error)

	GetUserByEmail(string) (entity.User, error)
	GetUser(int) (entity.User, error)

	UpdateUser(entity.User) (entity.User, error)
}
