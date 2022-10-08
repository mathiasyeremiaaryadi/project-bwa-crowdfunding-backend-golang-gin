package entity

import (
	"strings"
	"time"
)

type Campaign struct {
	ID               uint `gorm:"primaryKey"`
	UserId           uint
	Name             string `gorm:"varchar(30)"`
	ShortDescription string `gorm:"varchar(100)"`
	Description      string `gorm:"varchar(150)"`
	Perks            string `gorm:"varchar(50)"`
	BackerCount      int
	GoalAmount       int       `gorm:"varchar(255)"`
	CurrentAmount    int       `gorm:"varchar(255)"`
	Slug             string    `gorm:"varchar(50)"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
	CampaignImages   []CampaignImage
	User             User
}

type CampaignImage struct {
	ID         uint
	CampaignId uint
	FileName   string `gorm:"varchar(255)"`
	IsPrimary  int
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type GetCampaign struct {
	ID               uint   `json:"id"`
	UserId           uint   `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type GetCampaignDetail struct {
	ID               uint                     `json:"id"`
	UserId           uint                     `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageUrl         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             GetCampaignUserDetail    `json:"user"`
	Images           []GetCampaignImageDetail `json:"images"`
}

type GetCampaignUserDetail struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type GetCampaignImageDetail struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func getCampaignFormatter(campaign Campaign) GetCampaign {
	var getCampaignFormatter GetCampaign
	getCampaignFormatter.ID = campaign.ID
	getCampaignFormatter.UserId = campaign.UserId
	getCampaignFormatter.Name = campaign.Name
	getCampaignFormatter.ShortDescription = campaign.ShortDescription
	getCampaignFormatter.GoalAmount = campaign.GoalAmount
	getCampaignFormatter.CurrentAmount = campaign.CurrentAmount
	getCampaignFormatter.Slug = campaign.Slug
	getCampaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		getCampaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return getCampaignFormatter
}

func GetCampaignsFormatter(campaigns []Campaign) []GetCampaign {
	var getCampaignsFormatter []GetCampaign

	for _, campaign := range campaigns {
		getCampaignFormatter := getCampaignFormatter(campaign)
		getCampaignsFormatter = append(getCampaignsFormatter, getCampaignFormatter)
	}

	return getCampaignsFormatter
}

func GetCampaignDetailFormatter(campaign Campaign) GetCampaignDetail {
	var getCampaignDetailFormatter GetCampaignDetail
	getCampaignDetailFormatter.ID = campaign.ID
	getCampaignDetailFormatter.UserId = campaign.UserId
	getCampaignDetailFormatter.Name = campaign.Name
	getCampaignDetailFormatter.ShortDescription = campaign.ShortDescription
	getCampaignDetailFormatter.Description = campaign.Description
	getCampaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	getCampaignDetailFormatter.Slug = campaign.Slug
	getCampaignDetailFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		getCampaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))

	}
	getCampaignDetailFormatter.Perks = perks

	var user GetCampaignUserDetail
	user.Name = campaign.User.Name
	user.ImageUrl = campaign.User.AvatarFileName

	getCampaignDetailFormatter.User = user

	var images []GetCampaignImageDetail
	for _, campaignImage := range campaign.CampaignImages {
		var image GetCampaignImageDetail
		image.ImageUrl = campaignImage.FileName

		isPrimary := false
		if image.IsPrimary {
			isPrimary = true
		}
		image.IsPrimary = isPrimary
		images = append(images, image)
	}

	getCampaignDetailFormatter.Images = images

	return getCampaignDetailFormatter
}
