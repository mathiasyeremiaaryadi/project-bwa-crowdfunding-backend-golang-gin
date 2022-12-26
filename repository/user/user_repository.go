package userrepository

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

func (r *userRepository) RegisterUser(user entity.User) (entity.User, error) {
	if err := r.mysql.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.mysql.Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUser(id int) (entity.User, error) {
	var user entity.User

	if err := r.mysql.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	if err := r.mysql.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
