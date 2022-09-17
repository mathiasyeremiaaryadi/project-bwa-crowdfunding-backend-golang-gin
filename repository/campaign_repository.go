package repository

import (
	"service-campaign-startup/model/entity"

	"gorm.io/gorm"
)

type campaignRepository struct {
	mysql *gorm.DB
}

func NewCampaignRepository(mysql *gorm.DB) CampaignRepository {
	return &campaignRepository{
		mysql: mysql,
	}
}

func (repositories *campaignRepository) GetCampaigns() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	err := repositories.mysql.Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (repositories *campaignRepository) GetCampaignById(userId int) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	err := repositories.mysql.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = ?", 1).Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}
