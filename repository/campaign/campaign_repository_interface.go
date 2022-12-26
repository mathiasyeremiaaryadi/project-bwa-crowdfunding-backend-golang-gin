package campaignrepository

import "service-campaign-startup/model/entity"

type CampaignRepository interface {
	GetCampaigns() ([]entity.Campaign, error)
	GetCampaignByUserID(userID int) ([]entity.Campaign, error)
	GetCampaign(CampaignID int) (entity.Campaign, error)

	CreateCampaign(entity.Campaign) (entity.Campaign, error)
	CreateCampaignImage(entity.CampaignImage) error

	UpdateCampaign(entity.Campaign) (entity.Campaign, error)
	UpdateCampaignImageStatus(int) error
}
