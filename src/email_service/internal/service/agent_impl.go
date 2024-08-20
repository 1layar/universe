package service

import (
	"context"
	"crypto/tls"

	"github.com/1layar/merasa/backend/src/email_service/model"
	"gopkg.in/gomail.v2"
)

type emailAgent struct {
	account *accountService
}

// define const interface here
var _ Agent = (*emailAgent)(nil)

func NewEmailAgent(account *accountService) *emailAgent {
	return &emailAgent{
		account: account,
	}
}

func (a *emailAgent) SendEmail(ctx context.Context, message model.EmailMessage) error {
	var err error
	account := message.Account

	if account == nil {
		account, err = a.account.GetAccountByID(ctx, message.AccountID)

		if err != nil {
			return err
		}
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "merasa <noreply@kumerasa.com>")
	m.SetHeader("To", message.ToEmail)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/plain", message.TextBody)
	if message.HtmlBody != "" {
		m.AddAlternative("text/html", message.HtmlBody)
	}

	d := gomail.NewDialer(account.SMTPHost, account.SMTPPort, account.SMTPUsername, account.SMTPPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(m)
}
