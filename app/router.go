package app

import (
	"service-campaign-startup/config"
	"service-campaign-startup/delivery"
	campaigndelivery "service-campaign-startup/delivery/campaign"
	transactiondelivery "service-campaign-startup/delivery/transaction"
	userdelivery "service-campaign-startup/delivery/user"
	campaignrepository "service-campaign-startup/repository/campaign"
	transactionrepository "service-campaign-startup/repository/transaction"
	userrepository "service-campaign-startup/repository/user"
	campaignusecase "service-campaign-startup/usecase/campaign"
	paymentusecase "service-campaign-startup/usecase/payment"
	transactionusecase "service-campaign-startup/usecase/transaction"
	userusecase "service-campaign-startup/usecase/user"
	"service-campaign-startup/utils"

	"github.com/gin-gonic/gin"
)

func NewRoute(dependencies *config.DependencyFacade) *gin.Engine {

	jwtService := utils.NewJwtService()

	userRepository := userrepository.NewUserRepository(dependencies)
	userUseCase := userusecase.NewUserUseCase(userRepository, jwtService)
	userDelivery := userdelivery.NewUserDelivery(userUseCase)

	campaignRepository := campaignrepository.NewCampaignRepository(dependencies)
	campaignUseCase := campaignusecase.NewCampaignUseCase(campaignRepository)
	campaignDelivery := campaigndelivery.NewCampaignDelivery(campaignUseCase)

	transactionRepository := transactionrepository.NewTransactionRepository(dependencies)
	paymentUseCase := paymentusecase.NewPaymentUseCase(transactionRepository, campaignRepository)
	transactionUseCase := transactionusecase.NewTransactionUseCase(transactionRepository, campaignRepository, paymentUseCase)
	transactionDelivery := transactiondelivery.NewTransactionDelivery(transactionUseCase, paymentUseCase)

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
		apiRouter.PUT("/campaigns/:id", AuthMiddleware(userUseCase, jwtService), campaignDelivery.UpdateCampaign)

		apiRouter.POST("/campaigns", AuthMiddleware(userUseCase, jwtService), campaignDelivery.CreateCampaign)
		apiRouter.POST("/campaigns-image", AuthMiddleware(userUseCase, jwtService), campaignDelivery.CreateCampaignImage)

		apiRouter.GET("/transactions", AuthMiddleware(userUseCase, jwtService), transactionDelivery.GetTransactionsByUserID)
		apiRouter.POST("/transactions", AuthMiddleware(userUseCase, jwtService), transactionDelivery.CreateTransaction)
	}

	router.NoRoute(delivery.NoRoute)

	return router
}
