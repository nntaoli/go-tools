package email

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

type Sender struct {
	Nick     string
	Email    string
	Password string
	SmtpHost string
	SmtpPort int
}

func createEmailBody(sender Sender, to, subject, content string, isHtml bool) []byte {
	header := make(map[string]string)
	header["From"] = sender.Nick + " <" + sender.Email + ">"
	header["To"] = to
	header["Subject"] = subject

	if isHtml {
		header["Content-Type"] = "text/html; charset=UTF-8"
	}

	emailBody := ""
	for k, v := range header {
		emailBody += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	emailBody += "\r\n" + content

	return []byte(emailBody)
}

func SendEmail(sender Sender, to string, subject string, content string, isHtml bool) error {

	auth := &SmtpAuth{"", sender.Email, sender.Password, sender.SmtpHost}
	err := smtp.SendMail(fmt.Sprintf("%s:%d", sender.SmtpHost, sender.SmtpPort),
		auth, sender.Email, []string{to}, createEmailBody(sender, to, subject, content, isHtml))

	return err
}

func SendEmailWithTls(sender Sender, to string, subject string, content string, isHtml bool) error {
	tlsCon, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", sender.SmtpHost, sender.SmtpPort), nil)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(tlsCon, sender.SmtpHost)
	if err != nil {
		return err
	}
	defer c.Close()

	if ok, _ := c.Extension("AUTH"); ok {
		auth := &SmtpAuth{"", sender.Email, sender.Password, sender.SmtpHost}
		if err = c.Auth(auth); err != nil {
			log.Print("Error during AUTH", err)
			return err
		}
	}

	if err = c.Mail(sender.Email); err != nil {
		return err
	}

	if err = c.Rcpt(to); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(createEmailBody(sender, to, subject, content, isHtml))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()

}
