package transactionusecase

import "service-campaign-startup/model/dto"

type TransactionUseCase interface {
	GetTransactionsByCampaignID(dto.TransactionUri) *dto.ResponseContainer
}
