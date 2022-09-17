package dto

type UserRegisterRequest struct {
	Name       string `json:"NAME" binding:"required"`
	Occupation string `json:"OCCUPATION" binding:"required"`
	Email      string `json:"EMAIL" binding:"required,email"`
	Password   string `json:"PASSWORD" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"EMAIL" binding:"required,email"`
	Password string `json:"PASSWORD" binding:"required"`
}

type EmailCheckRequest struct {
	Email string `json:"EMAIL" binding:"required,email"`
}
