package mailclient

import (
	"fmt"

	"gopkg.in/sendgrid/sendgrid-go.v1"
)

type client struct {
	*sendgrid.SGClient
}

func (c *client) Send(m Mail) error {
	if sgMail, ok := m.(*sendgrid.SGMail); ok {
		return c.SGClient.Send(sgMail)
	}

	return fmt.Errorf("Trying to send non sendgrid mail type `%T`", m)
}

func (c *client) NewMail() Mail {
	return sendgrid.NewMail()
}

// NewSendGridClient creates send grid client
func NewSendGridClient(apiUser, apiKey string) Client {
	return &client{
		SGClient: sendgrid.NewSendGridClient(apiUser, apiKey),
	}
}

// NewSendGridClientWithAPIKey creates send grid client
func NewSendGridClientWithAPIKey(apiKey string) Client {
	return &client{
		SGClient: sendgrid.NewSendGridClientWithApiKey(apiKey),
	}
}
