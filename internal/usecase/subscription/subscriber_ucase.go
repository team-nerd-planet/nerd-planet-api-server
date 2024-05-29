package subscription

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/smtp"
	"path"
	"runtime"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/team-nerd-planet/api-server/infra/config"
	"github.com/team-nerd-planet/api-server/internal/entity"
)

type SubscriptionUsecase struct {
	subscriptionRepo entity.SubscriptionRepo
	conf             *config.Config
}

func NewSubscriptionUsecase(
	subscriptionRepo entity.SubscriptionRepo,
	config *config.Config,
) SubscriptionUsecase {
	return SubscriptionUsecase{
		subscriptionRepo: subscriptionRepo,
		conf:             config,
	}
}

func (su SubscriptionUsecase) Apply(subscription entity.Subscription) (*entity.Subscription, bool) {
	slog.Info("SubscriptionUsecase_Apply", "subscription", subscription)
	id, err := su.subscriptionRepo.ExistEmail(subscription.Email)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	token := emailToken{Subscription: subscription}
	slog.Info("Init Token", "token", token)

	token.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))

	if id == nil {
		token.Type = SUBSCRIBE
	} else {
		token.Type = RESUBSCRIBE
		token.Subscription.ID = uint(*id)
	}

	slog.Info("Token", "token", token)

	tokenStr, err := generateEmailToken(token, su.conf.Jwt.SecretKey)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	slog.Info("Generate Token", "str", tokenStr)
	slog.Info("Config", "su.conf.Smtp", su.conf.Smtp)

	if err := sendSubscribeMail(su.conf.Smtp.Host, su.conf.Smtp.Port, su.conf.Smtp.UserName, su.conf.Smtp.Password, subscription.Email, tokenStr); err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	return &subscription, true
}

func (su SubscriptionUsecase) CancelSubscription(email string) (*entity.Subscription, bool) {
	//이메일 조회
	//JWT 토큰 생성
	//이메일이 없으면 실패 반환
	//이메일이 있으면 DELETE 이메일 전송
	return nil, true
}

func (su SubscriptionUsecase) Subscribe(token string) (*entity.Subscription, bool) {
	var (
		emailToken   *emailToken
		err          error
		subscription *entity.Subscription
	)

	emailToken, err = verifyEmailToken(token, su.conf.Jwt.SecretKey)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	switch emailToken.Type {
	case SUBSCRIBE:
		subscription, err = su.subscriptionRepo.Create(emailToken.Subscription)
	case RESUBSCRIBE:
		subscription, err = su.subscriptionRepo.Update(int64(emailToken.Subscription.ID), emailToken.Subscription)
	case UNSUBSCRIBE:
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

func (su SubscriptionUsecase) Resubscribe(subscription entity.Subscription) (*entity.Subscription, bool) {
	return nil, false
}

func (su SubscriptionUsecase) Unsubscribe(email string) (*entity.Subscription, bool) {
	return nil, false
}

func sendSubscribeMail(host string, port int, userName, password, email, token string) error {
	data := struct {
		Token string
	}{
		Token: token,
	}

	slog.Info("sendSubscribeMail", "host", host, "port", port, "userName", userName, "password", password, "email", email, "token", token)

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

	auth := smtp.PlainAuth("", userName, password, host)
	from := userName
	to := []string{email}
	subject := "Subject: Nerd Planet 메일 인증\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + body.String())
	err = smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}

var (
	errTokenExpired            = errors.New("token has invalid claims: token is expired")
	errUnexpectedSigningMethod = errors.New("unexpected signing method: HMAC-SHA")
	errSignatureInvalid        = errors.New("token signature is invalid: signature is invalid")
)

type tokenType int

const (
	SUBSCRIBE tokenType = iota
	RESUBSCRIBE
	UNSUBSCRIBE
)

type emailToken struct {
	Type         tokenType           `json:"type"`
	Subscription entity.Subscription `json:"subscription"`
	jwt.RegisteredClaims
}

func generateEmailToken(token emailToken, secretKey string) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return newToken.SignedString([]byte(secretKey))
}

func verifyEmailToken(tokenString string, secretKey string) (*emailToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &emailToken{}, func(token *jwt.Token) (interface{}, error) {
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

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*emailToken)
	if !ok {
		return nil, errSignatureInvalid
	}

	return claims, nil
}
