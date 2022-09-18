package campaignusecase

import (
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	campaignrepository "service-campaign-startup/repository/campaign"
)

type campaignUseCase struct {
	campaignRepository campaignrepository.CampaignRepository
}

func NewCampaignUseCase(campaignRepository campaignrepository.CampaignRepository) CampaignUseCase {
	return &campaignUseCase{
		campaignRepository: campaignRepository,
	}
}

func (usecases *campaignUseCase) GetCampaigns(userId int) *dto.ResponseContainer {
	if userId != 0 {
		campaigns, err := usecases.campaignRepository.GetCampaignById(userId)
		if err != nil {
			err := map[string]interface{}{"ERROR": err.Error()}
			return dto.BuildResponse(
				"Database query error or database connection problem",
				"FAILED",
				http.StatusInternalServerError,
				err,
			)
		}

		if len(campaigns) == 0 {
			err := map[string]interface{}{"ERROR": err.Error()}
			return dto.BuildResponse(
				"User not found",
				"FAILED",
				http.StatusNotFound,
				err,
			)
		}

		getCampaigns := entity.GetCampaignsFormatter(campaigns)
		return dto.BuildResponse(
			"Users have retrieved successfully",
			"SUCCESS",
			http.StatusCreated,
			getCampaigns,
		)
	}

	campaigns, err := usecases.campaignRepository.GetCampaigns()
	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	if len(campaigns) == 0 {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"User not found",
			"FAILED",
			http.StatusNotFound,
			err,
		)
	}

	getCampaigns := entity.GetCampaignsFormatter(campaigns)
	return dto.BuildResponse(
		"Users have retrieved successfully",
		"SUCCESS",
		http.StatusCreated,
		getCampaigns,
	)
}
