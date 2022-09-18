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

func (usecases *userUseCase) RegisterUser(request dto.UserRegisterRequest) *dto.ResponseContainer {
	var user entity.User

	hashedPassword, errHash := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.MinCost,
	)
	if errHash != nil {
		err := map[string]interface{}{"ERROR": errHash.Error()}
		return dto.BuildResponse(
			"Password hash failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation
	user.PasswordHash = string(hashedPassword)
	user.Role = "user"

	user, err := usecases.userRepository.RegisterUser(user)
	if err != nil {
		err := map[string]interface{}{"ERROR": errHash.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	token, err := usecases.jwtService.GenerateToken(int(user.ID))
	if err != nil {
		err := map[string]interface{}{"ERROR": errHash.Error()}
		return dto.BuildResponse(
			"Token generation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
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

func (usecases *userUseCase) LoginUser(request dto.UserLoginRequest) *dto.ResponseContainer {
	var user entity.User
	var email string
	var password string

	email = request.Email
	password = request.Password

	user, err := usecases.userRepository.GetUserByEmail(email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Authentication failed",
			"FAILED",
			http.StatusUnauthorized,
			err,
		)
	}

	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	errVerify := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if errVerify != nil {
		err := map[string]interface{}{"ERROR": errVerify.Error()}
		return dto.BuildResponse(
			"Authentication failed",
			"FAILED",
			http.StatusUnauthorized,
			err,
		)
	}

	token, err := usecases.jwtService.GenerateToken(int(user.ID))
	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Token generation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
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

func (usecases *userUseCase) GetUserByEmail(request dto.EmailCheckRequest) (bool, error) {
	email := request.Email

	_, err := usecases.userRepository.GetUserByEmail(email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (usecases *userUseCase) GetUserById(id int) (entity.User, error) {
	user, err := usecases.userRepository.GetUserById(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, err
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func (usecases *userUseCase) SaveUserAvatar(id int, fileLocation string) *dto.ResponseContainer {
	user, err := usecases.userRepository.GetUserById(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"User not found",
			"FAILED",
			http.StatusNotFound,
			err,
		)
	}

	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	user.AvatarFileName = fileLocation
	user, err = usecases.userRepository.UpdateUser(user)
	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Avatar updation failed",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	return dto.BuildResponse(
		"Avatar updation success",
		"SUCCESS",
		http.StatusCreated,
		"Avatar has uploaded and updated successfully",
	)
}
