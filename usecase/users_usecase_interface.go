package usecase

import (
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
)

type UserUseCase interface {
	RegisterUser(dto.UserRegisterRequest) *dto.ResponseContainer
	LoginUser(dto.UserLoginRequest) *dto.ResponseContainer
	GetUserByEmail(dto.EmailCheckRequest) (bool, error)
	GetUserById(int) (entity.User, error)
	SaveUserAvatar(int, string) *dto.ResponseContainer
}
