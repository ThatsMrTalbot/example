package mailclient

import (
	"io"
	"net/mail"
	"time"
)

// Mail is an email instance
type Mail interface {
	AddAttachment(filename string, file io.Reader) error
	AddAttachmentFromStream(filename, file string)
	AddBcc(bcc string) error
	AddBccRecipient(recipient *mail.Address)
	AddBccRecipients(recipients []*mail.Address)
	AddBccs(bccs []string) error
	AddCc(cc string) error
	AddCcRecipient(recipient *mail.Address)
	AddCcRecipients(recipients []*mail.Address)
	AddCcs(ccs []string) error
	AddContentID(id, value string)
	AddHeader(header, value string)
	AddRecipient(recipient *mail.Address)
	AddRecipients(recipients []*mail.Address)
	AddTo(email string) error
	AddToName(name string)
	AddToNames(names []string)
	AddTos(emails []string) error
	HeadersString() (string, error)
	SetDate(date string)
	SetFrom(from string) error
	SetFromEmail(address *mail.Address)
	SetFromName(fromname string)
	SetHTML(html string)
	SetRFCDate(date time.Time)
	SetReplyTo(replyto string) error
	SetReplyToEmail(address *mail.Address)
	SetSubject(subject string)
	SetText(text string)
}

// Client to send emails
type Client interface {
	Send(Mail) error
	NewMail() Mail
}
