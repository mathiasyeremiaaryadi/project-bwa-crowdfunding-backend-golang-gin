package entity

import "time"

type Transaction struct {
	ID         uint `gorm:"primaryKey"`
	CampaignID int
	UserID     int
	Amount     int
	Status     string `gorm:"varchar(10)"`
	Code       string `gorm:"varchar(10)"`
	User       User
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type GetTransaction struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
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
