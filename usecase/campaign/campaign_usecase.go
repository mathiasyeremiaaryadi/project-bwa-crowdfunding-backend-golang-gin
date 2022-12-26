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

func (uc *campaignUseCase) GetCampaigns(userId int) *dto.ResponseContainer {
	var campaigns []entity.Campaign
	var err error

	if userId != 0 {
		campaigns, err = uc.campaignRepository.GetCampaignByUserId(userId)
	} else {
		campaigns, err = uc.campaignRepository.GetCampaigns()
	}

	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if len(campaigns) == 0 {
		return dto.BuildResponse(
			"User not found",
			"FAILED",
			http.StatusNotFound,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	formattedCampaigns := entity.GetCampaignsFormatter(campaigns)
	return dto.BuildResponse(
		"Campaigns have retrieved successfully",
		"SUCCESS",
		http.StatusOK,
		formattedCampaigns,
	)
}

func (uc *campaignUseCase) GetCampaignById(campaignUri dto.CampaignUri) *dto.ResponseContainer {
	campaign, err := uc.campaignRepository.GetCampaignById(campaignUri.ID)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if reflect.DeepEqual(campaign, entity.Campaign{}) {
		return dto.BuildResponse(
			"User not found",
			"FAILED",
			http.StatusNotFound,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	campaignDetail := entity.GetCampaignDetailFormatter(campaign)
	return dto.BuildResponse(
		"Campaign has retrieved successfully",
		"SUCCESS",
		http.StatusCreated,
		campaignDetail,
	)
}

func (uc *campaignUseCase) CreateCampaign(request dto.CampaignRequest) *dto.ResponseContainer {
	var campaign entity.Campaign
	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.GoalAmount = request.GoalAmount
	campaign.Perks = request.Perks
	campaign.UserId = request.User.ID

	slugName := fmt.Sprintf("%s %d", request.Name, request.User.ID)
	campaign.Slug = slug.Make(slugName)

	campaign, err := uc.campaignRepository.CreateCampaign(campaign)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	campaignDetail := entity.GetCampaignFormatter(campaign)
	return dto.BuildResponse(
		"Campaign has updated successfully",
		"SUCCESS",
		http.StatusCreated,
		campaignDetail,
	)
}

func (uc *campaignUseCase) UpdateCampaign(campaignId dto.CampaignUri, request dto.CampaignRequest) *dto.ResponseContainer {
	campaign, err := uc.campaignRepository.GetCampaignById(campaignId.ID)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if campaign.UserId != request.User.ID {
		return dto.BuildResponse(
			"Unauthorized",
			"FAILED",
			http.StatusUnauthorized,
			map[string]interface{}{"ERROR": "Not an owner of the campaign"},
		)
	}

	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.Perks = request.Perks
	campaign.GoalAmount = request.GoalAmount
	campaign.UserId = request.User.ID

	updatedCampaign, err := uc.campaignRepository.UpdateCampaign(campaign)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	return dto.BuildResponse(
		"Campaign has updated successfully",
		"SUCCESS",
		http.StatusOK,
		updatedCampaign,
	)
}

func (uc *campaignUseCase) CreateCampaignImage(request dto.CampaignImageRequest, fileLocation string) *dto.ResponseContainer {
	campaign, err := uc.campaignRepository.GetCampaignById(request.CampaignId)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	if campaign.UserId != request.User.ID {
		return dto.BuildResponse(
			"Unauthorized",
			"FAILED",
			http.StatusUnauthorized,
			map[string]interface{}{"ERROR": "Not an owner of the campaign"},
		)
	}

	tempIsPrimary := 0
	if request.IsPrimary {
		tempIsPrimary = 1

		err := uc.campaignRepository.UpdateCampaignImageStatus(request.CampaignId)
		if err != nil {
			return dto.BuildResponse(
				"Database query error or database connection problem",
				"FAILED",
				http.StatusInternalServerError,
				map[string]interface{}{"ERROR": err.Error()},
			)
		}
	}

	var campaignImage entity.CampaignImage
	campaignImage.CampaignId = uint(request.CampaignId)
	campaignImage.IsPrimary = tempIsPrimary
	campaignImage.FileName = fileLocation

	err = uc.campaignRepository.CreateCampaignImage(campaignImage)
	if err != nil {
		return dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			map[string]interface{}{"ERROR": err.Error()},
		)
	}

	return dto.BuildResponse(
		"Campaign image has created successfully",
		"SUCCESS",
		http.StatusOK,
		"Campaign image has uploaded and updated successfully",
	)
}
