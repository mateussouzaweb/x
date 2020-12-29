package mail

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/mateussouzaweb/x/env"
)

// Data type
type Data = map[string]interface{}

// Mail struct
type Mail struct {
	From     string
	FromName string
	To       string
	ReplyTo  string
	Subject  string
	Message  string
	Template string
	Data     Data
}

// FillMissing method
func (m *Mail) FillMissing() {

	if m.From == "" {
		m.From = env.Get("MAIL_FROM", "")
	}
	if m.FromName == "" {
		m.FromName = env.Get("MAIL_FROM_NAME", "")
	}
	if m.ReplyTo == "" {
		m.ReplyTo = m.From
	}

	if m.Template != "" {
		m.Data["appName"] = env.Get("APP_NAME", "")
		m.Data["year"] = time.Now().Format("2006")
	}

}

// DeliveryPlain method
func (m *Mail) DeliveryPlain() error {

	m.FillMissing()

	var content = fmt.Sprintf("To: %s\r\n"+
		"From: %s <%s>\r\n"+
		"Reply-To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", m.To, m.FromName, m.From, m.ReplyTo, m.Subject, m.Message)

	return connection.Delivery(m, []byte(content))
}

// DeliveryHTML method
func (m *Mail) DeliveryHTML() error {

	m.FillMissing()

	tmpl, err := template.ParseFiles(m.Template)

	if err != nil {
		return err
	}

	var content = fmt.Sprintf("To: %s\r\n"+
		"From: %s <%s>\r\n"+
		"Reply-To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-version: 1.1\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
		"\r\n", m.To, m.FromName, m.From, m.ReplyTo, m.Subject)

	var body bytes.Buffer

	body.Write([]byte(content))
	tmpl.Execute(&body, m.Data)

	return connection.Delivery(m, body.Bytes())
}
