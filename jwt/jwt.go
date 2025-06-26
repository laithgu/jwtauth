package jwt

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	config   Config
	redisCli *redis.Client
}

func New(cfg Config, redisCli *redis.Client) *JWT {
	return &JWT{
		config:   cfg,
		redisCli: redisCli,
	}
}

// 生成 Token
func (j *JWT) GenerateToken(userID int64, username, platform string) (string, error) {
	now := time.Now()
	expire := now.Add(time.Duration(j.config.ExpireHours) * time.Hour)

	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Platform: platform,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(j.config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// 解析 Token
func (j *JWT) ParseToken(ctx context.Context, tokenStr string) (*CustomClaims, error) {
	// 黑名单判断（同之前）
	if exist, _ := j.redisCli.Exists(ctx, j.config.RedisPrefix+tokenStr).Result(); exist == 1 {
		return nil, ErrTokenBlacklisted
	}

	// 解析 token
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.SecretKey), nil
	})
	if err != nil {
		return nil, ErrTokenInvalid
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, ErrTokenInvalid
	}

	return claims, nil
}

// Token 加入黑名单（登出）
func (j *JWT) InvalidateToken(ctx context.Context, tokenStr string, expireAt time.Time) error {
	ttl := time.Until(expireAt)
	if ttl <= 0 {
		ttl = time.Hour // fallback ttl
	}
	key := j.config.RedisPrefix + tokenStr
	return j.redisCli.Set(ctx, key, "1", ttl).Err()
}
