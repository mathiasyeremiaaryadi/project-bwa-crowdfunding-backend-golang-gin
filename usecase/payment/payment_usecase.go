package paymentusecase

import (
	"service-campaign-startup/model/entity"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type paymentUseCase struct {
}

func NewPaymentUseCase() PaymentUsecase {
	return &paymentUseCase{}
}

func (pu paymentUseCase) GetPaymentURL(transaction entity.Transaction, user entity.User) (string, error) {
	midtrans.ServerKey = ""
	midtrans.ClientKey = ""
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
