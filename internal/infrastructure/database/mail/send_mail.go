package mail

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail() error {
	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", "guira2002@gmail.com")
	m.SetHeader("Subject", "Hello! Emailln")
	m.SetBody("text/html", "Olá Usuário aqui é o <b>Emailln</b>!")

	return d.DialAndSend(m)
}
