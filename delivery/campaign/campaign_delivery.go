package campaigndelivery

import (
	"net/http"
	"strconv"

	campaignusecase "service-campaign-startup/usecase/campaign"

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
