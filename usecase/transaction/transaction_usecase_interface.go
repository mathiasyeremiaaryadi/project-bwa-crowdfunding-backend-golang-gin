package transactionusecase

import (
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
)

type TransactionUseCase interface {
	GetTransactionsByCampaignID(dto.TransactionUri) *dto.ResponseContainer
	GetTransactionsByUserID(int) *dto.ResponseContainer
	CreateTransaction(entity.TransactionCreated) *dto.ResponseContainer
}
