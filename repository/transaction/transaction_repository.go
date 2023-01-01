package transactionrepository

import (
	"service-campaign-startup/model/entity"

	"gorm.io/gorm"
)

type transactionrepository struct {
	mysql *gorm.DB
}

func NewTransactionRepository(mysql *gorm.DB) TransactionRepository {
	return &transactionrepository{
		mysql: mysql,
	}
}

func (r *transactionrepository) GetTransactionsByCampaignID(campaignID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.mysql.Preload("User").Where("campaign_id = ?", campaignID).Order("id DESC").Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}
