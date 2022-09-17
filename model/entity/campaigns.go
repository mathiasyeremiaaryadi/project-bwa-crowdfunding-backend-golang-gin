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
}

type CampaignImage struct {
	ID         int
	CampaignId int
	FileName   string `gorm:"varchar(255)"`
	IsPrimary  int
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
