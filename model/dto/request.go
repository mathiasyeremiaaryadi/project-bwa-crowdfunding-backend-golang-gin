package dto

type UserRegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type EmailCheckRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type Campaign struct {
	ID int `uri:"id" binding:"required"`
}
