package jwt

import "errors"

var (
	ErrTokenExpired     = errors.New("token 已过期")
	ErrTokenInvalid     = errors.New("无效的 token")
	ErrTokenBlacklisted = errors.New("token 已失效（黑名单）")
)
