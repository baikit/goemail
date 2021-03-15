package goemail

import (
	"net/smtp"
	"strings"
)

type EmailConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

func (email EmailConfig) SendEmail(from string, to string, subject string, body string) error {

	auth := smtp.PlainAuth("", email.User, email.Password, email.Host)

	contentType := "Content-Type: text/html; charset=UTF-8"

	content := email.formatHtmlBody(body)

	msg := []byte("To: " + to + "\r\nFrom: " + from + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + content)

	sendTo := strings.Split(to, ";")

	addr := email.Host + ":" + email.Port

	err := smtp.SendMail(addr, auth, from, sendTo, msg)

	return err

}

func (email EmailConfig) formatHtmlBody(msg string) string {

	return `<html>
		<body>
		<div>
		"` + msg + `"
		</div>
		</body>
		</html>`
}