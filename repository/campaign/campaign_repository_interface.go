package campaignrepository

import "service-campaign-startup/model/entity"

type CampaignRepository interface {
	GetCampaigns() ([]entity.Campaign, error)
	GetCampaignByUserId(userId int) ([]entity.Campaign, error)
	GetCampaignById(campaignId int) (entity.Campaign, error)
}
