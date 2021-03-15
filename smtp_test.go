package goemail

import (
	"fmt"
	"testing"
)

func TestEmailConfig_formatHtmlBody(t *testing.T) {

	config := EmailConfig{User: "@34", Password: "234", Host: "@34",Port: "587"}

	res := config.formatHtmlBody("@34")

	fmt.Println(res)

}

func TestEmailConfig_SendEmail(t *testing.T) {

	email := EmailConfig{User: "test", Password: "test", Host: "smtp.qq.com", Port: "587"}

	err := email.SendEmail("baikit", "985", "test", "激活码：007008")

	fmt.Println(err)

}
