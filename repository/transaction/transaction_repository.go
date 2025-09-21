package transactionrepository

import (
	"service-campaign-startup/config"
	"service-campaign-startup/model/entity"
)

type transactionrepository struct {
	dependencies *config.DependencyFacade
}

func NewTransactionRepository(dependencies *config.DependencyFacade) TransactionRepository {
	return &transactionrepository{
		dependencies: dependencies,
	}
}

func (r *transactionrepository) GetTransactionsByCampaignID(campaignID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.dependencies.MySQLDB.Debug().Preload("User").Where("campaign_id = ?", campaignID).Order("id DESC").Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionrepository) GetTransaction(transactionID int) (entity.Transaction, error) {
	var transaction entity.Transaction

	if err := r.dependencies.MySQLDB.Debug().Where("id = ?", transactionID).Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionrepository) GetTransactionsByUserID(userID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.dependencies.MySQLDB.Debug().Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionrepository) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	if err := r.dependencies.MySQLDB.Debug().Create(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionrepository) UpdateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	if err := r.dependencies.MySQLDB.Debug().Save(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}
