package domain

import "github.com/golang-jwt/jwt"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	UserName string `json:"userName"`
	RoleCode string `json:"roleCode"`
	IsActive *bool  `json:"isActive"`
	jwt.StandardClaims
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type GetUserRequest struct {
	UserName string `json:"userName"`
}

type UpdateUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
