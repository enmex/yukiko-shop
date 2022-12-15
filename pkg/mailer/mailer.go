package mailer

import (
	"github.com/enmex/smtp/config"
	s "github.com/enmex/smtp/sender"
)

var (
	gmailHost    = "smtp.gmail.com:"
	gmailAddress = "smtp.gmail.com:465"
)

type Config struct {
	User     string
	Password string
}

type Mailer struct {
	sender *s.Sender
}

func NewMailer(cfg *Config) *Mailer {
	providers := make(map[string]config.Provider, 1)
	providers["default"] = config.Provider{
		Credentials: config.Credentials{
			User:     cfg.User,
			Password: cfg.Password,
		},
		Delivery: config.Delivery{
			Host:    gmailHost,
			Address: gmailAddress,
		},
	}

	return &Mailer{
		sender: s.NewSender(config.Config{
			Mode:      config.SingleMode,
			Providers: providers,
		}),
	}
}

func (m *Mailer) SendMail(subject, message, sender, recipient string) error {
	return m.sender.Send(s.SendMailPayload{
		Provider:   "default",
		Subject:    subject,
		Message:    message,
		SenderMail: sender,
		Recipient:  recipient,
	})
}
