package entity

import "time"

type Campaign struct {
	ID               uint `gorm:"primaryKey"`
	UserId           int
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
}

type CampaignImage struct {
	ID         int
	CampaignId int
	FileName   string `gorm:"varchar(255)"`
	IsPrimary  int
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type GetCampaign struct {
	ID               uint   `json:"ID"`
	UserId           int    `json:"USER_ID"`
	Name             string `json:"NAME"`
	ShortDescription string `json:"SHORT_DESCRIPTION"`
	ImageUrl         string `json:"IMAGE_URL"`
	GoalAmount       int    `json:"GOAL_AMOUNT"`
	CurrentAmount    int    `json:"CURRENT_AMOUNT"`
}

func getCampaignFormatter(campaign Campaign) GetCampaign {
	var getCampaignFormatter GetCampaign
	getCampaignFormatter.ID = campaign.ID
	getCampaignFormatter.UserId = campaign.UserId
	getCampaignFormatter.Name = campaign.Name
	getCampaignFormatter.ShortDescription = campaign.ShortDescription
	getCampaignFormatter.GoalAmount = campaign.GoalAmount
	getCampaignFormatter.CurrentAmount = campaign.CurrentAmount

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
