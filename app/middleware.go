package app

import (
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/usecase"
	"service-campaign-startup/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(userUseCase usecase.UserUseCase, jwtService utils.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			err := map[string]interface{}{"ERROR": "Invalid authentication method"}
			response := dto.BuildResponse(
				"Authentication failed",
				"FAILED",
				http.StatusUnauthorized,
				err,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var token string
		authData := strings.Split(authHeader, " ")
		if len(authData) == 2 {
			token = authData[1]
		}

		validatedToken, err := jwtService.ValidateToken(token)
		if err != nil {
			err := map[string]interface{}{"ERROR": "Invalid token sign method"}
			response := dto.BuildResponse(
				"Authentication failed",
				"FAILED",
				http.StatusUnauthorized,
				err,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := validatedToken.Claims.(jwt.MapClaims)
		if !ok || !validatedToken.Valid {
			err := map[string]interface{}{"ERROR": "Invalid token claims"}
			response := dto.BuildResponse(
				"Authentication failed",
				"FAILED",
				http.StatusUnauthorized,
				err,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["USER_ID"].(float64))
		user, err := userUseCase.GetUserById(userId)
		if err != nil {
			err := map[string]interface{}{"ERROR": "User not found"}
			response := dto.BuildResponse(
				"Authentication failed",
				"FAILED",
				http.StatusUnauthorized,
				err,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("authenticatedUser", user)
	}

}
