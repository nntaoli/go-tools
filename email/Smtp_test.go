package email

import "testing"

func TestSendEmail(t *testing.T) {
	SendEmail(Sender{
		Nick:     "",
		Email:    "",
		Password: "",
		SmtpHost: "",
		SmtpPort: 25,
	},
		"xxx@163.com", "测试发送邮件", "测试发邮件", false)
}

func TestSendEmailWithTls(t *testing.T) {
	SendEmailWithTls(Sender{
		Nick:     "",
		Email:    "",
		Password: "",
		SmtpHost: "",
		SmtpPort: 465,
	},
		"xxx@163.com", "测试发邮件", "<b>测试发邮件</b>", true)
}
