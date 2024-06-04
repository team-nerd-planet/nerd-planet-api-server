package subscription

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"net/smtp"
	"path"
	"runtime"
	"time"

	"github.com/team-nerd-planet/api-server/infra/config"
	"github.com/team-nerd-planet/api-server/internal/entity"
	"github.com/team-nerd-planet/api-server/internal/usecase/token"
	"github.com/team-nerd-planet/api-server/internal/usecase/token/model"
)

type SubscriptionUsecase struct {
	subscriptionRepo entity.SubscriptionRepo
	tokenUcase       token.TokenUsecase
	conf             *config.Config
}

func NewSubscriptionUsecase(
	subscriptionRepo entity.SubscriptionRepo,
	tokenUsecase token.TokenUsecase,
	config *config.Config,
) SubscriptionUsecase {
	return SubscriptionUsecase{
		subscriptionRepo: subscriptionRepo,
		tokenUcase:       tokenUsecase,
		conf:             config,
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
		slog.Error(err.Error(), "error", err)
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
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	if err := sendSubscribeMail(su.conf.Smtp, name, subscription.Email, emailTokenStr); err != nil {
		slog.Error(err.Error(), "error", err)
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
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	return subscription, true
}

func sendSubscribeMail(smtpConf config.Smtp, name, email, token string) error {
	data := struct {
		Name  string
		Token string
	}{
		Name:  name,
		Token: token,
	}

	_, b, _, _ := runtime.Caller(0)
	configDirPath := path.Join(path.Dir(b))
	t, err := template.ParseFiles(fmt.Sprintf("%s/template/subscription.html", configDirPath))
	if err != nil {
		slog.Error(err.Error(), "error", err)
		return err
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		slog.Error(err.Error(), "error", err)
		return err
	}

	auth := smtp.PlainAuth("", smtpConf.UserName, smtpConf.Password, smtpConf.Host)
	from := smtpConf.UserName
	to := []string{email}
	subject := "Subject: Nerd Planet 메일 인증\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + body.String())
	err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpConf.Host, smtpConf.Port), auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
