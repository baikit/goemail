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

//SendEmail
//to 发送给谁 例子 发送多人用分号间隔 "111@qq.com;222@qq.com;333@qq.com"
//subject 邮件的主题
//body 邮件主要内容
func (email EmailConfig) SendEmail(to string, subject string, body string) error {

	auth := smtp.PlainAuth("", email.User, email.Password, email.Host)

	contentType := "Content-Type: text/html; charset=UTF-8"

	content := email.formatHtmlBody(body)

	msg := []byte("To: " + to + "\r\nFrom: " + email.User + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + content)

	sendTo := strings.Split(to, ";")

	addr := email.Host + ":" + email.Port

	err := smtp.SendMail(addr, auth, email.User, sendTo, msg)

	return err

}

func (email EmailConfig) formatHtmlBody(msg string) string {
	return `<html><body><div>`+ msg + `</div></body></html>`
}
