package userdelivery

import (
	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUserAvatar(c *gin.Context)
	GetAuthenticatedUser(c *gin.Context)
}
