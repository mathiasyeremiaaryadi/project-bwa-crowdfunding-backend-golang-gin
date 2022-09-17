package usecase

import "service-campaign-startup/model/dto"

type CampaignUseCase interface {
	GetCampaigns() *dto.ResponseContainer
}
