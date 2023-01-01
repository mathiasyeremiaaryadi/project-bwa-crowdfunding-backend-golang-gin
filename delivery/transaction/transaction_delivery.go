package transactiondelivery

import (
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	transactionusecase "service-campaign-startup/usecase/transaction"

	"github.com/gin-gonic/gin"
)

type transactionDelivery struct {
	transactionUseCase transactionusecase.TransactionUseCase
}

func NewTransactionDelivery(transactionUseCase transactionusecase.TransactionUseCase) TransactionDelivery {
	return &transactionDelivery{
		transactionUseCase: transactionUseCase,
	}
}

func (d *transactionDelivery) GetTransactionsByCampaignID(c *gin.Context) {
	var transactionUri dto.TransactionUri

	if err := c.ShouldBindUri(&transactionUri); err != nil {
		response := dto.BuildResponse(
			"Body request bind failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": err.Error()},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if authenticatedUser, ok := c.Value("authenticatedUser").(entity.User); ok {
		transactionUri.User = authenticatedUser
	} else {
		response := dto.BuildResponse(
			"Authentication failed",
			"FAILED",
			http.StatusUnauthorized,
			"not authenticated",
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := d.transactionUseCase.GetTransactionsByCampaignID(transactionUri)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
