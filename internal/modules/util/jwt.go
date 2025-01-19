package util

import (
	"libro-system-api/internal/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, AccountName string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	expiration := time.Second * time.Duration(cfg.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"accountName": AccountName,
		"expiredAt":   time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
