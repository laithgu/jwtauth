package jwt

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Platform string `json:"platform,omitempty"` // 可选字段，用于标记来源
	jwt.RegisteredClaims
}
