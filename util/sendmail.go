package util

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/go-gomail/gomail"
	"os"
)

func SendMail(from string, message string) error {
	email := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")
	anotherEmail := os.Getenv("ANOTHER_EMAIL_USERNAME")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email, anotherEmail)
	m.SetHeader("Subject", "New Message from Davillex.com!")
	m.SetBody("text/html", message)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, email, password)

	err := d.DialAndSend(m)

	return err
}
