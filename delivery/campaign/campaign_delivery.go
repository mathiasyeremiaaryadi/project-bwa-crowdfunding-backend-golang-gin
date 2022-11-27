package campaigndelivery

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
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
	var campaignUri dto.CampaignUri

	err := c.ShouldBindUri(&campaignUri)
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

	response := deliveries.campaignUseCase.GetCampaignById(campaignUri)
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

func (deliveries *campaignDelivery) CreateCampaign(c *gin.Context) {
	var request dto.CreateCampaignRequest

	err := c.ShouldBindJSON(&request)
	if err != nil && errors.Is(err, io.EOF) {
		err := map[string]interface{}{"ERROR": err.Error()}
		response := dto.BuildResponse(
			"Body request bind failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		errors := utils.ValidationFormatter(err)
		err := map[string]interface{}{"ERROR": errors}
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authenticatedUser := c.MustGet("authenticatedUser").(entity.User)
	request.User = authenticatedUser

	response := deliveries.campaignUseCase.CreateCampaign(request)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

}
