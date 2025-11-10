package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	UserID   string `json:"user_id"`
}
