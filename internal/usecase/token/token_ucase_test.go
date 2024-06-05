package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
	"github.com/team-nerd-planet/api-server/infra/config"
	"github.com/team-nerd-planet/api-server/internal/entity"
	"github.com/team-nerd-planet/api-server/internal/usecase/token/model"
)

func Test_GenerateAndVerifyEmailToken(t *testing.T) {
	tokenUsecase := NewTokenUsecase(config.Config{
		Jwt: config.Jwt{
			SecretKey: "test_key",
		},
	})

	name := "name"
	division := "division"
	token1 := model.EmailToken{
		Subscription: entity.Subscription{
			Email:                   "email",
			Name:                    &name,
			Division:                &division,
			Published:               time.Now(),
			PreferredCompanyArr:     pq.Int64Array{},
			PreferredCompanySizeArr: pq.Int64Array{},
			PreferredJobArr:         pq.Int64Array{},
			PreferredSkillArr:       pq.Int64Array{},
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	str, err := tokenUsecase.GenerateToken(token1)
	if err != nil {
		t.Error(err.Error())
		return

	}

	var token2 model.EmailToken
	err = tokenUsecase.VerifyToken(str, &token2)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if token1.Subscription.Email != token2.Subscription.Email ||
		*token1.Subscription.Name != *token2.Subscription.Name ||
		*token1.Subscription.Division != *token2.Subscription.Division ||
		token1.Subscription.Published.Compare(token2.Subscription.Published) != 0 {
		t.Errorf("not same")
		return
	}

	t.Log("PASS")
}
