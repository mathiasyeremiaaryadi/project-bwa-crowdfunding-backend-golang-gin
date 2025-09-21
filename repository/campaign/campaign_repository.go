package campaignrepository

import (
	"service-campaign-startup/config"
	"service-campaign-startup/model/entity"
)

type campaignRepository struct {
	dependencies *config.DependencyFacade
}

func NewCampaignRepository(dependencies *config.DependencyFacade) CampaignRepository {
	return &campaignRepository{
		dependencies: dependencies,
	}
}

func (r *campaignRepository) GetCampaigns() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	if err := r.dependencies.MySQLDB.Debug().Preload("CampaignImages", "campaign_images.is_primary = ?", 1).Find(&campaigns).Error; err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (r *campaignRepository) GetCampaignByUserID(userID int) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	if err := r.dependencies.MySQLDB.Debug().Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = ?", 1).Find(&campaigns).Error; err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (r *campaignRepository) GetCampaign(CampaignID int) (entity.Campaign, error) {
	var campaign entity.Campaign

	if err := r.dependencies.MySQLDB.Debug().Preload("User").Preload("CampaignImages").Where("id = ?", CampaignID).Find(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) CreateCampaign(campaign entity.Campaign) (entity.Campaign, error) {
	if err := r.dependencies.MySQLDB.Debug().Create(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) UpdateCampaign(campaign entity.Campaign) (entity.Campaign, error) {
	if err := r.dependencies.MySQLDB.Debug().Save(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) CreateCampaignImage(campaignImage entity.CampaignImage) error {
	if err := r.dependencies.MySQLDB.Debug().Create(&campaignImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *campaignRepository) UpdateCampaignImageStatus(CampaignID int) error {
	if err := r.dependencies.MySQLDB.Debug().Model(&entity.CampaignImage{}).Where("campaign_id = ?", CampaignID).Update("is_primary", false).Error; err != nil {
		return err
	}

	return nil
}
