package entity

import "time"

type Transaction struct {
	ID         uint `gorm:"primaryKey"`
	CampaignID int
	UserID     int
	Amount     int
	Status     string    `gorm:"varchar(10)"`
	Code       string    `gorm:"varchar(10)"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
