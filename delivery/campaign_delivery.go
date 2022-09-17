package delivery

import (
	"net/http"
	"service-campaign-startup/usecase"

	"github.com/gin-gonic/gin"
)

type campaignDelivery struct {
	campaignUseCase usecase.CampaignUseCase
}

func NewCampaignDelivery(campaignUseCase usecase.CampaignUseCase) CampaignDelivery {
	return &campaignDelivery{
		campaignUseCase: campaignUseCase,
	}
}

func (deliveries *campaignDelivery) GetCampaigns(c *gin.Context) {
	response := deliveries.campaignUseCase.GetCampaigns()
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
