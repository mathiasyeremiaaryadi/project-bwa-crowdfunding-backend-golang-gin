package transactiondelivery

import (
	"github.com/gin-gonic/gin"
)

type TransactionDelivery interface {
	GetTransactionsByCampaignID(c *gin.Context)
	GetTransactionsByUserID(c *gin.Context)
}
