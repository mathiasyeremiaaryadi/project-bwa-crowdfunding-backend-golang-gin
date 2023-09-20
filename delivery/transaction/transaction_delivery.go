package transactiondelivery

import (
	"errors"
	"io"
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	transactionusecase "service-campaign-startup/usecase/transaction"
	"service-campaign-startup/utils"

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
	if response.Meta.Code != http.StatusOK {
		c.AbortWithStatusJSON(response.Meta.Code, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (d *transactionDelivery) GetTransactionsByUserID(c *gin.Context) {
	var userID int

	if authenticatedUser, ok := c.Value("authenticatedUser").(entity.User); ok {
		userID = int(authenticatedUser.ID)
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

	response := d.transactionUseCase.GetTransactionsByUserID(userID)
	if response.Meta.Code != http.StatusOK {
		c.AbortWithStatusJSON(response.Meta.Code, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (d *transactionDelivery) CreateTransaction(c *gin.Context) {
	var transactionCreated entity.TransactionCreated

	if err := c.ShouldBindJSON(&transactionCreated); err != nil {
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

	if authenticatedUser, ok := c.Value("authenticatedUser").(entity.User); ok {
		transactionCreated.User = authenticatedUser
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

	response := d.transactionUseCase.CreateTransaction(transactionCreated)
	if response.Meta.Code != http.StatusOK {
		c.AbortWithStatusJSON(response.Meta.Code, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
