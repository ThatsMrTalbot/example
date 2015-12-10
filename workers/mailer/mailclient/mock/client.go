package mock

import (
	"fmt"

	"github.com/ThatsMrTalbot/example/workers/mailer/mailclient"
)

// Client is the mock emailer client
type Client struct {
	Sent []*Mail
}

// Send mail client
func (c *Client) Send(m mailclient.Mail) error {
	if mockM, ok := m.(*Mail); ok {
		c.Sent = append(c.Sent, mockM)
		return nil
	}

	return fmt.Errorf("Trying to send non mock mail type `%T`", m)
}

// NewMail creates client
func (c *Client) NewMail() mailclient.Mail {
	return &Mail{}
}

// NewMockEmailClient Create mock email client
func NewMockEmailClient() *Client {
	return &Client{}
}
