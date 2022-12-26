package userusecase

import (
	"errors"
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	userrepository "service-campaign-startup/repository/user"
	"service-campaign-startup/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUseCase struct {
	userRepository userrepository.UserRepository
	jwtService     utils.JWT
}

func NewUserUseCase(userRepository userrepository.UserRepository, jwtService utils.JWT) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (uc *userUseCase) RegisterUser(request dto.UserRegisterRequest) *dto.ResponseContainer {
	var user entity.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return dto.BuildResponse(
			"Password hash failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation
	user.PasswordHash = string(hashedPassword)
	user.Role = "user"

	user, err = uc.userRepository.RegisterUser(user)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	token, err := uc.jwtService.GenerateToken(int(user.ID))
	if err != nil {
		return dto.BuildResponse(
			"Token generation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	createdUser := entity.UserCreatedFormatter(user, token)
	return dto.BuildResponse(
		"Registration success",
		"SUCCESS",
		http.StatusCreated,
		createdUser,
	)
}

func (uc *userUseCase) LoginUser(request dto.UserLoginRequest) *dto.ResponseContainer {
	user, err := uc.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.BuildResponse(
				"Authentication failed",
				"FAILED",
				http.StatusUnauthorized,
				map[string]interface{}{"ERROR": err.Error()},
			)
		}

		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		return dto.BuildResponse(
			"Authentication failed",
			"FAILED",
			http.StatusUnauthorized,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	token, err := uc.jwtService.GenerateToken(int(user.ID))
	if err != nil {
		return dto.BuildResponse(
			"Token generation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	authenticatedUser := entity.UserCreatedFormatter(user, token)
	return dto.BuildResponse(
		"Authentication success",
		"SUCCESS",
		http.StatusOK,
		authenticatedUser,
	)
}

func (uc *userUseCase) GetUserByEmail(request dto.EmailCheckRequest) (bool, error) {
	_, err := uc.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (uc *userUseCase) GetUserById(id int) (entity.User, error) {
	user, err := uc.userRepository.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *userUseCase) CreateUserAvatar(id int, fileLocation string) *dto.ResponseContainer {
	user, err := uc.userRepository.GetUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.BuildResponse(
				"User not found",
				"FAILED",
				http.StatusNotFound,
				map[string]interface{}{"ERROR": err.Error()},
			)
		}

		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	user.AvatarFileName = fileLocation
	user, err = uc.userRepository.UpdateUser(user)
	if err != nil {
		return dto.BuildResponse(
			"Avatar updation failed",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	return dto.BuildResponse(
		"Avatar updation success",
		"SUCCESS",
		http.StatusCreated,
		"Avatar has uploaded and updated successfully",
	)
}
