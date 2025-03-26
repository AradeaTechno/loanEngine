package helpers

import (
	"amarthaloan/config"
	"bytes"
	"text/template"

	"gopkg.in/mail.v2"
)

type EmailData struct {
	Email     string `json:"email"`
	TypeEmail string `json:"typeEmail"`
	AppName   string `json:"appName"`
	Subject   string `json:"subject"`
	Link      string `json:"link"`
	TeamEmail string `json:"teamEmail"`
}

func SendEmail(emailData EmailData) error {
	emailConf := config.EmailConfig()
	// confApp := config.AppConfig()

	// Parse html template
	tmpl, err := template.ParseFiles("email_template/" + emailData.TypeEmail + ".html")
	if err != nil {
		return err
	}

	// Buffer to hold executed template data
	var body bytes.Buffer

	// Inject dynamic data into the template and execute it
	if err := tmpl.Execute(&body, emailData); err != nil {
		return err
	}

	// Set up gomail
	m := mail.NewMessage()
	m.SetHeader("From", emailConf.SMTP_USER)
	m.SetHeader("To", emailData.Email)
	m.SetHeader("Subject", emailData.Subject)

	// Set the mail body as rendered HTML template
	m.SetBody("text/html", body.String())

	// Set up SMTP Dialer
	d := mail.NewDialer(emailConf.SMTP_HOST, emailConf.SMTP_PORT, emailConf.SMTP_USER, emailConf.SMTP_PASSWORD)
	// Send email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
