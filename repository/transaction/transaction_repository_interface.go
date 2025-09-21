package transactionrepository

import "service-campaign-startup/model/entity"

type TransactionRepository interface {
	GetTransactionsByCampaignID(int) ([]entity.Transaction, error)
	GetTransactionsByUserID(int) ([]entity.Transaction, error)
	GetTransaction(int) (entity.Transaction, error)
	CreateTransaction(entity.Transaction) (entity.Transaction, error)
	UpdateTransaction(entity.Transaction) (entity.Transaction, error)
}
