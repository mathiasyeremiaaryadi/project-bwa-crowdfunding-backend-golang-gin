package campaignusecase

import (
	"fmt"
	"net/http"
	"reflect"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	campaignrepository "service-campaign-startup/repository/campaign"

	"github.com/gosimple/slug"
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
		campaigns, err := usecases.campaignRepository.GetCampaignByUserId(userId)
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
			err := map[string]interface{}{"ERROR": "Not Found"}
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
		"Campaigns have retrieved successfully",
		"SUCCESS",
		http.StatusOK,
		getCampaigns,
	)
}

func (usecases *campaignUseCase) GetCampaignById(campaignUri dto.CampaignUri) *dto.ResponseContainer {

	campaign, err := usecases.campaignRepository.GetCampaignById(campaignUri.ID)
	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	if reflect.DeepEqual(campaign, entity.Campaign{}) {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"User not found",
			"FAILED",
			http.StatusNotFound,
			err,
		)
	}

	getCampaignDetail := entity.GetCampaignDetailFormatter(campaign)
	return dto.BuildResponse(
		"Campaign has retrieved successfully",
		"SUCCESS",
		http.StatusCreated,
		getCampaignDetail,
	)
}

func (usecases *campaignUseCase) CreateCampaign(request dto.CampaignRequest) *dto.ResponseContainer {
	var campaign entity.Campaign
	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.GoalAmount = request.GoalAmount
	campaign.Perks = request.Perks
	campaign.UserId = request.User.ID

	slugName := fmt.Sprintf("%s %d", request.Name, request.User.ID)
	campaign.Slug = slug.Make(slugName)

	campaign, err := usecases.campaignRepository.CreateCampaign(campaign)
	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	getCampaignDetail := entity.GetCampaignFormatter(campaign)
	return dto.BuildResponse(
		"Campaign has updated successfully",
		"SUCCESS",
		http.StatusCreated,
		getCampaignDetail,
	)
}

func (usecases *campaignUseCase) UpdateCampaign(campaignId dto.CampaignUri, request dto.CampaignRequest) *dto.ResponseContainer {
	campaign, err := usecases.campaignRepository.GetCampaignById(campaignId.ID)
	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	if campaign.UserId != request.User.ID {
		err := map[string]interface{}{"ERROR": "Not an owner of the campaign"}
		return dto.BuildResponse(
			"Unauthorized",
			"FAILED",
			http.StatusUnauthorized,
			err,
		)
	}

	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.Perks = request.Perks
	campaign.GoalAmount = request.GoalAmount
	campaign.UserId = request.User.ID

	updatedCampaign, err := usecases.campaignRepository.UpdateCampaign(campaign)
	if err != nil {
		err := map[string]interface{}{"ERROR": err.Error()}
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
	}

	return dto.BuildResponse(
		"Campaign has updated successfully",
		"SUCCESS",
		http.StatusOK,
		updatedCampaign,
	)
}
