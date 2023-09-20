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

func (r *transactionrepository) GetTransactionsByUserID(userID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.mysql.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionrepository) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	if err := r.mysql.Create(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}
