package mail

import (
	"bytes"
	"fmt"
	"text/template"
)

// From struct
type From struct {
	Email string
	Name  string
}

// Data type
type Data = map[string]interface{}

// Mail struct
type Mail struct {
	From     From
	To       string
	ReplyTo  string
	Subject  string
	HTML     bool
	Message  string
	Template string
	Data     Data
}

// Fill method
func (m *Mail) Fill() {

	if m.From.Name == "" {
		m.From.Name = _config.From.Name
	}
	if m.From.Email == "" {
		m.From.Email = _config.From.Email
	}
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

	var tmpl *template.Template
	var err error

	if m.HTML {
		tmpl, err = template.ParseFiles(m.Template)
	} else {
		tmpl, err = template.New("plain").Parse(m.Message)
	}

	if err != nil {
		return err
	}

	content := fmt.Sprintf("To: %s\r\n"+
		"From: %s <%s>\r\n"+
		"Reply-To: %s\r\n"+
		"Subject: %s\r\n",
		m.To, m.From.Name, m.From.Email, m.ReplyTo, m.Subject)

	if m.HTML {
		content += "MIME-version: 1.1\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n"
	} else {
		content += "\r\n"
	}

	var body bytes.Buffer

	body.Write([]byte(content))
	tmpl.Execute(&body, m.Data)

	return _config.SMTP.Delivery(m, body.Bytes())
}
