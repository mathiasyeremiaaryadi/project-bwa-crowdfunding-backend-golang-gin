package repository

import "service-campaign-startup/model/entity"

type CampaignRepository interface {
	GetCampaigns() ([]entity.Campaign, error)
	GetCampaignById(userId int) ([]entity.Campaign, error)
}
