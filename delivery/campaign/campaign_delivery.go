package campaigndelivery

import (
	"net/http"
	"strconv"

	"service-campaign-startup/model/dto"
	campaignusecase "service-campaign-startup/usecase/campaign"
	"service-campaign-startup/utils"

	"github.com/gin-gonic/gin"
)

type campaignDelivery struct {
	campaignUseCase campaignusecase.CampaignUseCase
}

func NewCampaignDelivery(campaignUseCase campaignusecase.CampaignUseCase) CampaignDelivery {
	return &campaignDelivery{
		campaignUseCase: campaignUseCase,
	}
}

func (deliveries *campaignDelivery) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query(("user_id")))

	response := deliveries.campaignUseCase.GetCampaigns(userId)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if response.Meta.Code == http.StatusNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	c.JSON(http.StatusOK, response)

}

func (deliveries *campaignDelivery) GetCampaignById(c *gin.Context) {
	var campaign dto.Campaign

	err := c.ShouldBindUri(&campaign)
	if err != nil {
		errors := utils.ValidationFormatter(err)
		err := map[string]interface{}{"ERROR": errors}
		response := dto.BuildResponse(
			"URI validation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := deliveries.campaignUseCase.GetCampaignById(campaign)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if response.Meta.Code == http.StatusNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	c.JSON(http.StatusOK, response)

}
