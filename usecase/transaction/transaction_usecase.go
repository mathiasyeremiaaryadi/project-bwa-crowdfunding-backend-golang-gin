package transactionusecase

import (
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	campaignrepository "service-campaign-startup/repository/campaign"
	transactionrepository "service-campaign-startup/repository/transaction"
	paymentusecase "service-campaign-startup/usecase/payment"
)

type transactionUseCase struct {
	transactionrepository transactionrepository.TransactionRepository
	campaignRepository    campaignrepository.CampaignRepository
	paymentUseCase        paymentusecase.PaymentUsecase
}

func NewTransactionUseCase(transactionrepository transactionrepository.TransactionRepository, campaignRepository campaignrepository.CampaignRepository, paymentUseCase paymentusecase.PaymentUsecase) TransactionUseCase {
	return &transactionUseCase{
		transactionrepository: transactionrepository,
		campaignRepository:    campaignRepository,
		paymentUseCase:        paymentUseCase,
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

	if len(transactions) == 0 {
		return dto.BuildResponse(
			"Transactions not found",
			"FAILED",
			http.StatusNotFound,
			map[string]interface{}{"ERROR": "not found"},
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

func (uc *transactionUseCase) GetTransactionsByUserID(userID int) *dto.ResponseContainer {
	transactions, err := uc.transactionrepository.GetTransactionsByUserID(userID)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if len(transactions) == 0 {
		return dto.BuildResponse(
			"Transactions not found",
			"FAILED",
			http.StatusNotFound,
			map[string]interface{}{"ERROR": "not found"},
		)
	}

	formattedTransactions := entity.GetTransactionsByIDFormatter(transactions)
	return dto.BuildResponse(
		"Transactions have retrieved successfully",
		"SUCCESS",
		http.StatusOK,
		formattedTransactions,
	)
}

func (uc *transactionUseCase) CreateTransaction(transactionCreated entity.TransactionCreated) *dto.ResponseContainer {
	transaction := entity.Transaction{}
	transaction.CampaignID = transactionCreated.CampaignID
	transaction.Amount = transactionCreated.Amount
	transaction.UserID = transactionCreated.User.ID
	transaction.Status = "pending"

	savedTransaction, err := uc.transactionrepository.CreateTransaction(transaction)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	paymentURL, err := uc.paymentUseCase.GetPaymentURL(savedTransaction, savedTransaction.User)
	if err != nil {
		return dto.BuildResponse(
			"Cannot retrieve payment URL",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	savedTransaction.PaymentURL = paymentURL
	savedTransaction, err = uc.transactionrepository.UpdateTransaction(savedTransaction)
	if err != nil {
		return dto.BuildResponse(
			"Failed to update transaction",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	return dto.BuildResponse(
		"Transactions have saved successfully",
		"SUCCESS",
		http.StatusCreated,
		entity.GetTransactionPaymentFormatter(savedTransaction),
	)
}
