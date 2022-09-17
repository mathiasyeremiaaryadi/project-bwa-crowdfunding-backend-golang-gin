package repository

import (
	"service-campaign-startup/model/entity"
)

type UserRepository interface {
	RegisterUser(entity.User) (entity.User, error)
	GetUserByEmail(string) (entity.User, error)
	GetUserById(int) (entity.User, error)
	UpdateUser(entity.User) (entity.User, error)
}
