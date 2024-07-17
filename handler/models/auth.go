package models

type LoginRequest struct {
	Username string `json:"username" bingding:"required"`
	Password string `json:"password" bingding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" bingding:"required"`
	Password string `json:"password" bingding:"required"`
	Age      int    `json:"age" bingding:"required"`
	Name     string `json:"name" bingding:"required"`
}
