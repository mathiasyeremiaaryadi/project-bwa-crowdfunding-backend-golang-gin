package entity

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"varchar(30)"`
	Occupation     string    `gorm:"varchar(30)"`
	Email          string    `gorm:"varchar(30)"`
	PasswordHash   string    `gorm:"varchar(255)"`
	AvatarFileName string    `gorm:"varchar(30)" `
	Role           string    `gorm:"varchar(10)"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

type UserCreated struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func UserCreatedFormatter(user User, token string) UserCreated {
	return UserCreated{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
}
