package entity

import "time"

type Transaction struct {
	ID         uint `gorm:"primaryKey"`
	CampaignID int
	UserID     uint
	Amount     int
	Status     string `gorm:"varchar(10)"`
	Code       string `gorm:"varchar(10)"`
	PaymentURL string
	User       User
	Campaign   Campaign
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type GetTransaction struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type GetTransactionByUserID struct {
	ID        uint                `json:"id"`
	Amount    int                 `json:"amount"`
	Status    string              `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	Campaign  TransactionCampaign `json:"campaign"`
}

type TransactionCampaign struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type TransactionCreated struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       User
}

type TransactionPayment struct {
	ID         uint      `json:"id"`
	CampaignID int       `json:"campaign_id"`
	UserID     uint      `json:"user_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentURL string    `json:"payment_url"`
	CreatedAt  time.Time `json:"created_at"`
}

func GetTransactionFormatter(transaction Transaction) GetTransaction {
	return GetTransaction{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
}

func GetTransactionsFormatter(transactions []Transaction) []GetTransaction {
	if len(transactions) == 0 {
		return []GetTransaction{}
	}

	var getTransactionsFormatter []GetTransaction

	for _, transaction := range transactions {
		getTransactionFormatter := GetTransactionFormatter(transaction)
		getTransactionsFormatter = append(getTransactionsFormatter, getTransactionFormatter)
	}

	return getTransactionsFormatter
}

func GetTransactionByUserIDFormatter(transaction Transaction) GetTransactionByUserID {
	var getTransactionByUserID GetTransactionByUserID
	getTransactionByUserID.ID = transaction.ID
	getTransactionByUserID.Amount = transaction.Amount
	getTransactionByUserID.Status = transaction.Status
	getTransactionByUserID.CreatedAt = transaction.CreatedAt

	var transactionCampaign TransactionCampaign
	transactionCampaign.Name = transaction.Campaign.Name
	transactionCampaign.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		transactionCampaign.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	getTransactionByUserID.Campaign = transactionCampaign

	return getTransactionByUserID
}

func GetTransactionsByIDFormatter(transactions []Transaction) []GetTransactionByUserID {
	if len(transactions) == 0 {
		return []GetTransactionByUserID{}
	}

	var getTransactionsByUserID []GetTransactionByUserID

	for _, transaction := range transactions {
		getTransactionFormatter := GetTransactionByUserIDFormatter(transaction)
		getTransactionsByUserID = append(getTransactionsByUserID, getTransactionFormatter)
	}

	return getTransactionsByUserID
}

func GetTransactionPaymentFormatter(transaction Transaction) TransactionPayment {
	var tranasctionPayment TransactionPayment
	tranasctionPayment.ID = transaction.ID
	tranasctionPayment.CampaignID = transaction.CampaignID
	tranasctionPayment.UserID = transaction.UserID
	tranasctionPayment.Status = transaction.Status
	tranasctionPayment.Amount = transaction.Amount
	tranasctionPayment.Code = transaction.Code
	tranasctionPayment.PaymentURL = transaction.PaymentURL
	tranasctionPayment.CreatedAt = transaction.CreatedAt

	return tranasctionPayment
}
