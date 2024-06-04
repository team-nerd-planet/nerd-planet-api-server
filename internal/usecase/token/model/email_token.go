package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/team-nerd-planet/api-server/internal/entity"
)

type TokenType int

const (
	SUBSCRIBE TokenType = iota
	RESUBSCRIBE
	UNSUBSCRIBE
)

type EmailToken struct {
	TokenType    TokenType           `json:"token_type"`
	Subscription entity.Subscription `json:"subscription"`
	jwt.RegisteredClaims
}

func NewEmailToken(tokenType TokenType, subscription entity.Subscription) EmailToken {
	return EmailToken{
		TokenType:    tokenType,
		Subscription: subscription,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
}
