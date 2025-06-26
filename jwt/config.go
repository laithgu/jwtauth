package jwt

type Config struct {
	SecretKey   string // JWT 秘钥
	ExpireHours int64  // Token 过期小时数
	RedisPrefix string // Redis 前缀，用于黑名单或唯一 Token 存储
	StrictMode  bool   // 是否启用 Redis 校验模式（互踢、强一致性等）
}
