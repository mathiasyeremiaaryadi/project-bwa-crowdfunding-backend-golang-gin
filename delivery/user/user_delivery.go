package userdelivery

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	userusecase "service-campaign-startup/usecase/user"
	"service-campaign-startup/utils"

	"github.com/gin-gonic/gin"
)

type userDelivery struct {
	userUseCase userusecase.UserUseCase
}

func NewUserDelivery(userUseCase userusecase.UserUseCase) UserDelivery {
	return &userDelivery{
		userUseCase: userUseCase,
	}
}

func (d *userDelivery) RegisterUser(c *gin.Context) {
	var request dto.UserRegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		if errors.Is(err, io.EOF) {
			response := dto.BuildResponse(
				"Body request bind failed",
				"FAILED",
				http.StatusBadRequest,
				map[string]interface{}{"errors": err.Error()},
			)

			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		errors := utils.ValidationFormatter(err)
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": errors},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := d.userUseCase.RegisterUser(request)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (d *userDelivery) LoginUser(c *gin.Context) {
	var request dto.UserLoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		if errors.Is(err, io.EOF) {
			response := dto.BuildResponse(
				"Body request bind failed",
				"FAILED",
				http.StatusBadRequest,
				map[string]interface{}{"errors": err.Error()},
			)

			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		errors := utils.ValidationFormatter(err)
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": errors},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := d.userUseCase.LoginUser(request)
	if response.Meta.Code != http.StatusInternalServerError {
		c.AbortWithStatusJSON(response.Meta.Code, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (d *userDelivery) GetUserByEmail(c *gin.Context) {
	var request dto.EmailCheckRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		if errors.Is(err, io.EOF) {
			response := dto.BuildResponse(
				"Body request bind failed",
				"FAILED",
				http.StatusBadRequest,
				map[string]interface{}{"errors": err.Error()},
			)

			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		errors := utils.ValidationFormatter(err)
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": errors},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isEmailExist, err := d.userUseCase.GetUserByEmail(request)
	if err != nil {
		response := dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"errors": err.Error()},
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if isEmailExist {
		response := dto.BuildResponse(
			"Email registration faled",
			"FAILED",
			http.StatusUnprocessableEntity,
			"Email already registered",
		)

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := dto.BuildResponse(
		"Email has retrieved successfully",
		"SUCCESS",
		http.StatusOK,
		"Email is available",
	)

	c.JSON(http.StatusOK, response)
}

func (d *userDelivery) GetUser(c *gin.Context) {
	var request dto.EmailCheckRequest

	err := c.ShouldBindJSON(&request)
	if err != nil && errors.Is(err, io.EOF) {
		err := map[string]interface{}{"errors": err.Error()}
		response := dto.BuildResponse(
			"Body request bind failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		errors := utils.ValidationFormatter(err)
		err := map[string]interface{}{"errors": errors}
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isEmailRegistered, err := d.userUseCase.GetUserByEmail(request)
	if err != nil {
		err := map[string]interface{}{"errors": err.Error()}
		response := dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if isEmailRegistered {
		response := dto.BuildResponse(
			"Email registration faled",
			"FAILED",
			http.StatusUnprocessableEntity,
			"Email already registered",
		)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := dto.BuildResponse(
		"Email has retrieved successfully",
		"SUCCESS",
		http.StatusOK,
		"Email is available",
	)
	c.JSON(http.StatusOK, response)
}

func (d *userDelivery) CreateUserAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		response := dto.BuildResponse(
			"Avatar upload failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": err.Error()},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authenticatedUser, ok := c.Value("authenticatedUser").(entity.User)
	if !ok {
		response := dto.BuildResponse(
			"Authentication failed",
			"FAILED",
			http.StatusUnauthorized,
			"not authenticated",
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	fullPath := fmt.Sprintf("images/%d-%s", authenticatedUser.ID, file.Filename)
	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		response := dto.BuildResponse(
			"Avatar upload failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": err.Error()},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := d.userUseCase.CreateUserAvatar(1, fullPath)
	if response.Meta.Code != http.StatusOK {
		c.AbortWithStatusJSON(response.Meta.Code, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
