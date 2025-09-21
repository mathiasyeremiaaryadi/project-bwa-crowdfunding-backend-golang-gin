package userrepository

import (
	"service-campaign-startup/config"
	"service-campaign-startup/model/entity"
)

type userRepository struct {
	dependencies *config.DependencyFacade
}

func NewUserRepository(dependencies *config.DependencyFacade) UserRepository {
	return &userRepository{
		dependencies: dependencies,
	}
}

func (r *userRepository) RegisterUser(user entity.User) (entity.User, error) {
	if err := r.dependencies.MySQLDB.Debug().Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.dependencies.MySQLDB.Debug().Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUser(id int) (entity.User, error) {
	var user entity.User

	if err := r.dependencies.MySQLDB.Debug().First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	if err := r.dependencies.MySQLDB.Debug().Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
