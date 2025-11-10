package model

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	LoginKey string `json:"loginKey" binding:"required"`
	Password string `json:"password" binding:"required"`
}
