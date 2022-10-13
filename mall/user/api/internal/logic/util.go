package logic

import (
	"go-zero-demo2/mall/pkg/jwtx"
	"go-zero-demo2/mall/user/api/internal/config"
	"time"
)

func getJwtToken(config config.Config, userId string) (string, error) {
	return jwtx.GetToken(config.Auth.AccessSecret, time.Now().Unix(), config.Auth.AccessExpire, userId)
}
