package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"hrm/config"
	"hrm/internal/domain"
)

func createSession(userName string) (string, *time.Time, error) {
	t := time.Now()
	expireTime := t.Add(24 * 3 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.Claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	})
	key := []byte(config.GetInstance().SecretKey)
	accessToken, err := token.SignedString(key)

	if err != nil {
		return "", &expireTime, err
	}

	return accessToken, &expireTime, nil
}
