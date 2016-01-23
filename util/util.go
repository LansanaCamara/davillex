package util

import (
	"github.com/go-gomail/gomail"
	_ "github.com/joho/godotenv/autoload"

	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// Regular expression helper function. Takes an expression
// and something to match it with -- both string types.
//
// Returns a boolean.
func MatchRegexp(expression string, str string) bool {
	re := regexp.MustCompile(expression)
	return re.Match([]byte(str))
}

// Checks whether a string is empty.
//
// returns a boolean.
func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// Takes a http.ResponseWriter type, filename as string, and
// data as interface type and renders HTML with the data.
//
// For example: util.RenderView(w, "index", args)
func RenderView(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles("public/views/" + filename + ".html")
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

// Send mail. Takes a "from" email address, and a message
// as arguments. Sends the email to two predefined addresses
// that are sent in the .env file.
//
// Returns an error in the case of the mail not being sent,
// or nil if no error occurred.
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
