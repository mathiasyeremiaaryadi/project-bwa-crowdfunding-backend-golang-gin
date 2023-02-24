package transactionrepository

import "service-campaign-startup/model/entity"

type TransactionRepository interface {
	GetTransactionsByCampaignID(int) ([]entity.Transaction, error)
	GetTransactionsByUserID(int) ([]entity.Transaction, error)
}
