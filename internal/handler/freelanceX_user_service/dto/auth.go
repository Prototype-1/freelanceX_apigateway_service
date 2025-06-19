package dto

type RegisterRequest struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
    Role     string `json:"role" binding:"required"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type SelectRoleRequest struct {
	UserId string `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"`
}


type OAuthRequest struct {
	Code     string `json:"code" binding:"required"`
	Provider string `json:"provider" binding:"required"`
}