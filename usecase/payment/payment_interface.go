package paymentusecase

import (
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
)

type PaymentUsecase interface {
	GetPaymentURL(transaction entity.Transaction, user entity.User) (string, error)
	ProcessPayment(transactionNotification entity.TransactionNotification) *dto.ResponseContainer
}
