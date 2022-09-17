package app

import (
	"service-campaign-startup/delivery"
	"service-campaign-startup/repository"
	"service-campaign-startup/usecase"
	"service-campaign-startup/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoute(mysql *gorm.DB) *gin.Engine {

	jwtService := utils.NewJwtService()

	userRepository := repository.NewUserRepository(mysql)
	userUseCase := usecase.NewUserUseCase(userRepository, jwtService)
	userDelivery := delivery.NewUserDelivery(userUseCase)

	campaignRepository := repository.NewCampaignRepository(mysql)
	campaignUseCase := usecase.NewCampaignUseCase(campaignRepository)
	campaignDelivery := delivery.NewCampaignDelivery(campaignUseCase)

	router := gin.Default()

	apiRouter := router.Group("/api/v1")
	{
		apiRouter.POST("/login", userDelivery.LoginUser)
		apiRouter.POST("/register", userDelivery.RegisterUser)
		apiRouter.POST("/email_checkers", userDelivery.GetUserByEmail)
		apiRouter.POST("/avatars", AuthMiddleware(userUseCase, jwtService), userDelivery.SaveUserAvatar)

		apiRouter.GET("/campaigns", campaignDelivery.GetCampaigns)
	}

	router.NoRoute(delivery.NoRoute)

	return router
}
