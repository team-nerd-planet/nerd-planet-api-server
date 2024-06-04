package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/team-nerd-planet/api-server/infra/config"
)

var (
	errTokenExpired            = errors.New("token has invalid claims: token is expired")
	errUnexpectedSigningMethod = errors.New("unexpected signing method: HMAC-SHA")
	errSignatureInvalid        = errors.New("token signature is invalid: signature is invalid")
)

type TokenUsecase struct {
	conf *config.Config
}

func NewTokenUsecase(conf *config.Config) TokenUsecase {
	return TokenUsecase{
		conf: conf,
	}
}

func (tu TokenUsecase) GenerateToken(claims jwt.Claims) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newToken.SignedString([]byte(tu.conf.Jwt.SecretKey))
}

func (tu TokenUsecase) VerifyToken(tokenString string, claims jwt.Claims) (err error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnexpectedSigningMethod
		}
		expiration, err := token.Claims.GetExpirationTime()
		if err != nil {
			return nil, err
		}

		if expiration.Time.Unix() < time.Now().Unix() {
			return nil, errTokenExpired
		}

		return []byte(tu.conf.Jwt.SecretKey), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errSignatureInvalid
	}

	return nil
}
