package campaigndelivery

import (
	"github.com/gin-gonic/gin"
)

type CampaignDelivery interface {
	GetCampaigns(c *gin.Context)
	GetCampaignById(c *gin.Context)

	CreateCampaign(c *gin.Context)
	CreateCampaignImage(c *gin.Context)

	UpdateCampaign(c *gin.Context)
}
