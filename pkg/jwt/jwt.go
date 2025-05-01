package jwt

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	redis "github.com/Prototype-1/freelanceX_apigateway_service/pkg/redis"
	"github.com/Prototype-1/freelanceX_apigateway_service/config"
	"log"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func ParseAccessToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}

func ValidateSession(sessionID, userID string) bool {
	ctx := context.Background()
	storedUserID, err := redis.RedisClient.Get(ctx, "session:"+sessionID).Result()
	if err != nil {
		log.Println("Redis session error:", err)
		return false
	}
	return storedUserID == userID
}
