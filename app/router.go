package app

import (
	"service-campaign-startup/delivery"
	campaigndelivery "service-campaign-startup/delivery/campaign"
	transactiondelivery "service-campaign-startup/delivery/transaction"
	userdelivery "service-campaign-startup/delivery/user"
	campaignrepository "service-campaign-startup/repository/campaign"
	transactionrepository "service-campaign-startup/repository/transaction"
	userrepository "service-campaign-startup/repository/user"
	campaignusecase "service-campaign-startup/usecase/campaign"
	transactionusecase "service-campaign-startup/usecase/transaction"
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

	transactionRepository := transactionrepository.NewTransactionRepository(mysql)
	transactionUseCase := transactionusecase.NewTransactionUseCase(transactionRepository, campaignRepository)
	transactionDelivery := transactiondelivery.NewTransactionDelivery(transactionUseCase)

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
		apiRouter.GET("/campaigns/:id/transactions", AuthMiddleware(userUseCase, jwtService), transactionDelivery.GetTransactionsByCampaignID)
		apiRouter.GET("/transactions", AuthMiddleware(userUseCase, jwtService), transactionDelivery.GetTransactionsByUserID)

		apiRouter.POST("/campaigns", AuthMiddleware(userUseCase, jwtService), campaignDelivery.CreateCampaign)
		apiRouter.POST("/campaigns-image", AuthMiddleware(userUseCase, jwtService), campaignDelivery.CreateCampaignImage)
		apiRouter.POST("/transactions", AuthMiddleware(userUseCase, jwtService), transactionDelivery.CreateTransaction)

		apiRouter.PUT("/campaigns/:id", AuthMiddleware(userUseCase, jwtService), campaignDelivery.UpdateCampaign)

	}

	router.NoRoute(delivery.NoRoute)

	return router
}
