package campaignrepository

import "service-campaign-startup/model/entity"

type CampaignRepository interface {
	GetCampaigns() ([]entity.Campaign, error)
	GetCampaignByUserId(userId int) ([]entity.Campaign, error)
	GetCampaignById(campaignId int) (entity.Campaign, error)

	CreateCampaign(entity.Campaign) (entity.Campaign, error)
	CreateCampaignImage(entity.CampaignImage) error

	UpdateCampaign(entity.Campaign) (entity.Campaign, error)
	UpdateCampaignImageStatus(int) error
}
