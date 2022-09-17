package repository

import (
	"service-campaign-startup/model/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	mysql *gorm.DB
}

func NewUserRepository(mysql *gorm.DB) UserRepository {
	return &userRepository{
		mysql: mysql,
	}
}

func (repositories *userRepository) RegisterUser(user entity.User) (entity.User, error) {
	err := repositories.mysql.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repositories *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User

	err := repositories.mysql.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repositories *userRepository) GetUserById(id int) (entity.User, error) {
	var user entity.User

	err := repositories.mysql.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repositories *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	err := repositories.mysql.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
