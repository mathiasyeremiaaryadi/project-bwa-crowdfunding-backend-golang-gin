package usecase

import (
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/repository"
)

type campaignUseCase struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignUseCase(campaignRepository repository.CampaignRepository) CampaignUseCase {
	return &campaignUseCase{
		campaignRepository: campaignRepository,
	}
}

func (usecases *campaignUseCase) GetCampaigns() *dto.ResponseContainer {
	users, err := usecases.campaignRepository.GetCampaigns()
	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	if len(users) == 0 {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"User not found",
			"FAILED",
			http.StatusNotFound,
			err,
		)
	}

	return dto.BuildResponse(
		"Users have retrieved successfully",
		"SUCCESS",
		http.StatusCreated,
		users,
	)
}
