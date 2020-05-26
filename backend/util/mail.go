package util

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

// mail send email
func mail(mailTo []string, subject string, body string) error {
	// TODO complete the  password of smtp mail
	mailConn := map[string]string{
		"user": "shoppie_ltd@163.com",
		"pass": "",
		"host": "smtp.163.com",
		"port": "465",
	}
	port, _ := strconv.Atoi(mailConn["port"])

	mailMessage := gomail.NewMessage()
	mailMessage.SetHeader("From", mailMessage.FormatAddress(mailConn["user"], "shoppie"))
	mailMessage.SetHeader("To", mailTo...)
	mailMessage.SetHeader("Subject", subject)
	mailMessage.SetBody("text/html", body)

	md := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := md.DialAndSend(mailMessage)
	return err
}

// SendMail send email
func SendMail() {
	// email receivers
	mailTo := []string{
		"1048143755@qq.com",
	}
	// email subject
	subject := "Greeting from golang mail client"
	// email body message
	body := "This is a mail sent by golang client used to check connectivity. If you see the email, the connection is successful."

	err := mail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return
	}
	fmt.Println("send successfully")
}
