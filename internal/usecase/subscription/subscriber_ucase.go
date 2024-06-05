package subscription

import (
	"log/slog"
	"time"

	"github.com/team-nerd-planet/api-server/internal/entity"
	"github.com/team-nerd-planet/api-server/internal/usecase/smtp"
	"github.com/team-nerd-planet/api-server/internal/usecase/token"
	"github.com/team-nerd-planet/api-server/internal/usecase/token/model"
)

type SubscriptionUsecase struct {
	subscriptionRepo entity.SubscriptionRepo
	tokenUcase       token.TokenUsecase
	smtpUcase        smtp.SmtpUsecase
}

func NewSubscriptionUsecase(
	subscriptionRepo entity.SubscriptionRepo,
	tokenUsecase token.TokenUsecase,
	smtpUsecase smtp.SmtpUsecase,
) SubscriptionUsecase {
	return SubscriptionUsecase{
		subscriptionRepo: subscriptionRepo,
		tokenUcase:       tokenUsecase,
		smtpUcase:        smtpUsecase,
	}
}

func (su SubscriptionUsecase) Apply(subscription entity.Subscription) (*entity.Subscription, bool) {
	var (
		emailToken    model.EmailToken
		emailTokenStr string
		name          string
	)

	subscription.Published = time.Now()

	if subscription.Name != nil {
		name = *subscription.Name
	} else {
		name = subscription.Email
	}

	id, err := su.subscriptionRepo.ExistEmail(subscription.Email)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	if id == nil {
		emailToken = model.NewEmailToken(model.SUBSCRIBE, subscription)
	} else {
		subscription.ID = uint(*id)
		emailToken = model.NewEmailToken(model.RESUBSCRIBE, subscription)
	}

	emailTokenStr, err = su.tokenUcase.GenerateToken(emailToken)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	if !su.smtpUcase.SendSubscriptionMail(name, subscription.Email, emailTokenStr) {
		return nil, false
	}

	return &subscription, true
}

func (su SubscriptionUsecase) Approve(token string) (*entity.Subscription, bool) {
	var (
		emailToken   model.EmailToken
		err          error
		subscription *entity.Subscription
	)

	if err := su.tokenUcase.VerifyToken(token, &emailToken); err != nil {
		slog.Warn(err.Error())
		return nil, true
	}

	switch emailToken.TokenType {
	case model.SUBSCRIBE:
		subscription, err = su.subscriptionRepo.Create(emailToken.Subscription)
	case model.RESUBSCRIBE:
		subscription, err = su.subscriptionRepo.Update(int64(emailToken.Subscription.ID), emailToken.Subscription)
	case model.UNSUBSCRIBE:
		subscription, err = su.subscriptionRepo.Delete(int64(emailToken.Subscription.ID))
	default:
		return nil, false
	}

	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return subscription, true
}
