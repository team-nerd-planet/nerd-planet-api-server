package smtp

import (
	"bytes"
	"log/slog"
	"net/smtp"
	"path"
	"runtime"
	"strconv"
	"text/template"

	"github.com/team-nerd-planet/api-server/infra/config"
)

const (
	TEMPLATE_DIR_NAME              = "template"
	SUBSCRIPTION_TEMPLATE_FILENAME = "subscription.html"
)

type SmtpUsecase struct {
	smtpConf config.Smtp
}

func NewSmtpUsecase(conf config.Config) SmtpUsecase {
	return SmtpUsecase{
		smtpConf: conf.Smtp,
	}
}

func (su SmtpUsecase) SendSubscriptionMail(name, email, token string) bool {
	data := struct {
		Name  string
		Token string
	}{
		Name:  name,
		Token: token,
	}

	if !su.sendMail(data, SUBSCRIPTION_TEMPLATE_FILENAME, "Nerd Planet 메일 인증", email) {
		slog.Error("failed to send the subscription mail")
		return false
	}

	return true
}

func (su SmtpUsecase) sendMail(data any, fileName, subject, email string) bool {
	_, b, _, _ := runtime.Caller(0)

	template, err := template.ParseFiles(path.Join(path.Join(path.Dir(b)), TEMPLATE_DIR_NAME, fileName))
	if err != nil {
		slog.Error(err.Error())
		return false
	}

	var body bytes.Buffer
	if err := template.Execute(&body, data); err != nil {
		slog.Error(err.Error())
		return false
	}

	addr := su.smtpConf.Host + ":" + strconv.Itoa(su.smtpConf.Port)
	auth := smtp.PlainAuth("", su.smtpConf.UserName, su.smtpConf.Password, su.smtpConf.Host)
	from := su.smtpConf.UserName
	to := []string{email}
	msg := []byte("Subject: " + subject + "\n" + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + body.String())

	if err = smtp.SendMail(addr, auth, from, to, msg); err != nil {
		return false
	}

	return true
}
