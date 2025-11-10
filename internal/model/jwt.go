package model

import (
	"time"
)

type Config struct {
	Issuer         string
	SecretKey      string
	ExpirationTime time.Duration
}

var DefaultJWTConfig = Config{
	Issuer:         "Login_And_Register_Demo",
	SecretKey:      "MySuperSecureJWTKey@2025!DoNotShare!",
	ExpirationTime: time.Hour * 24 * 7,
}
