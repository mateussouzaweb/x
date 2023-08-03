package mail

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// Address struct
type Address struct {
	Reference string
	Name      string
	Email     string
}

// Data type
type Data = map[string]any

// Mail struct
type Mail struct {
	From    Address
	To      string
	ReplyTo string
	Subject string
	Message string
	Data    Data
}

// Fill method
func (m *Mail) Fill() {

	if m.ReplyTo == "" {
		m.ReplyTo = m.From.Email
	}

	for key, value := range _config.Data {
		m.Data[key] = value
	}

}

// Delivery method
func (m *Mail) Delivery() error {

	m.Fill()

	var body bytes.Buffer

	content := m.Message
	isHTML := strings.Contains(content, "</p>")

	body.WriteString(fmt.Sprintf("To: %s\r\n"+
		"From: %s <%s>\r\n"+
		"Reply-To: %s\r\n"+
		"Subject: %s\r\n",
		m.To, m.From.Name, m.From.Email, m.ReplyTo, m.Subject))

	if isHTML {
		body.WriteString("MIME-version: 1.1\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n")
	} else {
		body.WriteString("\r\n")
	}

	message, err := template.New("plain").Parse(content)

	if err != nil {
		return err
	}

	message.Execute(&body, m.Data)

	return _config.SMTP.Delivery(m, body.Bytes())
}
