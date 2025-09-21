package paymentusecase

import (
	"net/http"
	"reflect"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	campaignrepository "service-campaign-startup/repository/campaign"
	transactionrepository "service-campaign-startup/repository/transaction"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type paymentUseCase struct {
	transactionRepository transactionrepository.TransactionRepository
	campaignRepository    campaignrepository.CampaignRepository
}

func NewPaymentUseCase(transactionRepository transactionrepository.TransactionRepository, campaignRepository campaignrepository.CampaignRepository) PaymentUsecase {
	return &paymentUseCase{
		transactionRepository: transactionRepository,
		campaignRepository:    campaignRepository,
	}
}

func (pu paymentUseCase) GetPaymentURL(transaction entity.Transaction, user entity.User) (string, error) {
	midtrans.ServerKey = "Mid-server-_MQzecDq22Aw1vZ2eHXBkDSq"
	midtrans.ClientKey = "Mid-client-si5Pu9NOYCVpMrkU"
	midtrans.Environment = midtrans.Sandbox

	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(transaction.ID)),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapResponse, err := snap.CreateTransaction(snapRequest)
	if err != nil {
		return "", err
	}

	return snapResponse.RedirectURL, nil
}

func (pu paymentUseCase) ProcessPayment(transactionNotification entity.TransactionNotification) *dto.ResponseContainer {
	transactionId, _ := strconv.Atoi(transactionNotification.OrderID)

	transaction, err := pu.transactionRepository.GetTransaction(transactionId)
	if reflect.DeepEqual(transaction, entity.Transaction{}) {
		return dto.BuildResponse(
			"Transaction not found",
			"FAILED",
			http.StatusNotFound,
			map[string]interface{}{"ERROR": "Transaction not found"},
		)
	}

	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if transactionNotification.PaymentType == "credit_card" && transactionNotification.TransactionStatus == "capture" && transactionNotification.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if transactionNotification.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if transactionNotification.TransactionStatus == "deny" || transactionNotification.TransactionStatus == "expired" || transactionNotification.TransactionStatus == "cancel" {
		transaction.Status = "cancel"
	}

	updatedTransaction, err := pu.transactionRepository.UpdateTransaction(transaction)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	campaign, err := pu.campaignRepository.GetCampaign(updatedTransaction.CampaignID)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := pu.campaignRepository.UpdateCampaign(campaign)
		if err != nil {
			return dto.BuildResponse(
				"Database query error or database connection problem",
				"FAILED",
				http.StatusInternalServerError,
				map[string]interface{}{"ERROR": err.Error()},
			)
		}
	}

	return dto.BuildResponse(
		"Payment has processed successfully",
		"SUCCESS",
		http.StatusOK,
		nil,
	)
}
