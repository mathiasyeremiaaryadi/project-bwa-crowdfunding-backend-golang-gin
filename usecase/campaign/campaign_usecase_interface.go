package campaignusecase

import "service-campaign-startup/model/dto"

type CampaignUseCase interface {
	GetCampaigns(userId int) *dto.ResponseContainer
	GetCampaignById(dto.CampaignUri) *dto.ResponseContainer
	CreateCampaign(dto.CampaignRequest) *dto.ResponseContainer
	UpdateCampaign(dto.CampaignUri, dto.CampaignRequest) *dto.ResponseContainer
}
