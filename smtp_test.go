package goemail

import (
	"fmt"
	"testing"
)

func TestEmailConfig_formatHtmlBody(t *testing.T) {

	config := EmailConfig{User: "@34", Password: "234", Host: "@34", Port: "587"}

	res := config.formatHtmlBody("@34")

	fmt.Println(res)

}

func TestEmailConfig_SendEmail(t *testing.T) {

	email := EmailConfig{User: "xx@xx.com", Password: "xxx", Host: "smtp.qq.com", Port: "587"}

	err := email.SendEmail("xx@qq.com", "test", "激活码：007008")

	fmt.Println(err)

}
