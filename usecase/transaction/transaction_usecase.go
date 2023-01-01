package transactionusecase

import (
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	campaignrepository "service-campaign-startup/repository/campaign"
	transactionrepository "service-campaign-startup/repository/transaction"
)

type transactionUseCase struct {
	transactionrepository transactionrepository.TransactionRepository
	campaignRepository    campaignrepository.CampaignRepository
}

func NewTransactionUseCase(transactionrepository transactionrepository.TransactionRepository, campaignRepository campaignrepository.CampaignRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionrepository: transactionrepository,
		campaignRepository:    campaignRepository,
	}
}

func (uc *transactionUseCase) GetTransactionsByCampaignID(transactionUri dto.TransactionUri) *dto.ResponseContainer {
	campaign, err := uc.campaignRepository.GetCampaign(transactionUri.ID)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if campaign.User.ID != transactionUri.User.ID {
		return dto.BuildResponse(
			"Unauthorized",
			"FAILED",
			http.StatusUnauthorized,
			map[string]interface{}{"ERROR": "Not an owner of the campaign"},
		)
	}

	transactions, err := uc.transactionrepository.GetTransactionsByCampaignID(transactionUri.ID)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	formattedTransactions := entity.GetTransactionsFormatter(transactions)
	return dto.BuildResponse(
		"Transactions have retrieved successfully",
		"SUCCESS",
		http.StatusOK,
		formattedTransactions,
	)
}
