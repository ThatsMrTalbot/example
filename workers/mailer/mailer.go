package main

import (
	"github.com/ThatsMrTalbot/example"
	"github.com/ThatsMrTalbot/example/workers/mailer/mailclient"
)

// Mailer wrapper
type Mailer struct {
	Config *example.Config `inject:""`

	client mailclient.Client
}

// Open SendGrid account
func (m *Mailer) Open() error {
	m.client = mailclient.NewSendGridClient(m.Config.SendGrid.Username, m.Config.SendGrid.Key)

	return nil
}

// Client returns sendgrid client
func (m *Mailer) Client() mailclient.Client {
	return m.client
}
