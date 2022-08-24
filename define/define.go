package define

import "github.com/golang-jwt/jwt/v4"

type SystemConfig struct {
	Port  string `json:"port"`  // 端口
	Entry string `json:"entry"` // 入口地址
}

type UserBasic struct {
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 密码
}

type UserClaim struct {
	jwt.RegisteredClaims
}

var (
	Key = []byte("op-panel")
	PID int
)
