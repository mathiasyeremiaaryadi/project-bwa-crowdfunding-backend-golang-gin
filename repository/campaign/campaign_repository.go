package campaignrepository

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

func (r *campaignRepository) GetCampaigns() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	if err := r.mysql.Preload("CampaignImages", "campaign_images.is_primary = ?", 1).Find(&campaigns).Error; err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (r *campaignRepository) GetCampaignByUserID(userID int) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	if err := r.mysql.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = ?", 1).Find(&campaigns).Error; err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (r *campaignRepository) GetCampaign(CampaignID int) (entity.Campaign, error) {
	var campaign entity.Campaign

	if err := r.mysql.Preload("User").Preload("CampaignImages").Where("id = ?", CampaignID).Find(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) CreateCampaign(campaign entity.Campaign) (entity.Campaign, error) {
	if err := r.mysql.Create(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) UpdateCampaign(campaign entity.Campaign) (entity.Campaign, error) {
	if err := r.mysql.Save(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) CreateCampaignImage(campaignImage entity.CampaignImage) error {
	if err := r.mysql.Create(&campaignImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *campaignRepository) UpdateCampaignImageStatus(CampaignID int) error {
	if err := r.mysql.Model(&entity.CampaignImage{}).Where("campaign_id = ?", CampaignID).Update("is_primary", false).Error; err != nil {
		return err
	}

	return nil
}
