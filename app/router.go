package app

import (
	"service-campaign-startup/delivery"
	campaigndelivery "service-campaign-startup/delivery/campaign"
	userdelivery "service-campaign-startup/delivery/user"
	campaignrepository "service-campaign-startup/repository/campaign"
	userrepository "service-campaign-startup/repository/user"
	campaignusecase "service-campaign-startup/usecase/campaign"
	userusecase "service-campaign-startup/usecase/user"
	"service-campaign-startup/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoute(mysql *gorm.DB) *gin.Engine {

	jwtService := utils.NewJwtService()

	userRepository := userrepository.NewUserRepository(mysql)
	userUseCase := userusecase.NewUserUseCase(userRepository, jwtService)
	userDelivery := userdelivery.NewUserDelivery(userUseCase)

	campaignRepository := campaignrepository.NewCampaignRepository(mysql)
	campaignUseCase := campaignusecase.NewCampaignUseCase(campaignRepository)
	campaignDelivery := campaigndelivery.NewCampaignDelivery(campaignUseCase)

	router := gin.Default()
	router.Static("/images", "./images")

	apiRouter := router.Group("/api/v1")
	{
		apiRouter.POST("/login", userDelivery.LoginUser)
		apiRouter.POST("/register", userDelivery.RegisterUser)
		apiRouter.POST("/email_checkers", userDelivery.GetUserByEmail)
		apiRouter.POST("/avatars", AuthMiddleware(userUseCase, jwtService), userDelivery.CreateUserAvatar)

		apiRouter.GET("/campaigns", campaignDelivery.GetCampaigns)
		apiRouter.GET("/campaigns/:id", campaignDelivery.GetCampaign)

		apiRouter.POST("/campaigns", AuthMiddleware(userUseCase, jwtService), campaignDelivery.CreateCampaign)
		apiRouter.POST("/campaigns-image", AuthMiddleware(userUseCase, jwtService), campaignDelivery.CreateCampaignImage)

		apiRouter.PUT("/campaigns/:id", AuthMiddleware(userUseCase, jwtService), campaignDelivery.UpdateCampaign)
	}

	router.NoRoute(delivery.NoRoute)

	return router
}
