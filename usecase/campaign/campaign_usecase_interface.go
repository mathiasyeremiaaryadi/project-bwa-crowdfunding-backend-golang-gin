package campaignusecase

import "service-campaign-startup/model/dto"

type CampaignUseCase interface {
	GetCampaigns(userID int) *dto.ResponseContainer
	GetCampaign(dto.CampaignUri) *dto.ResponseContainer

	CreateCampaign(dto.CampaignRequest) *dto.ResponseContainer
	CreateCampaignImage(dto.CampaignImageRequest, string) *dto.ResponseContainer

	UpdateCampaign(dto.CampaignUri, dto.CampaignRequest) *dto.ResponseContainer
}
