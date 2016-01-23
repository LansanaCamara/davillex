package util

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"

	_ "github.com/joho/godotenv/autoload"
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
func SendMail(w http.ResponseWriter, r *http.Request, body string) error {
	email1 := os.Getenv("EMAIL_ONE")
	email2 := os.Getenv("EMAIL_TWO")

	ctx := appengine.NewContext(r)

	msg := &mail.Message{
		Sender:  "Davillex.com Contact Form <noreply@davillex.com>",
		To:      []string{email1, email2},
		Subject: "New Message from Davillex.com",
		HTMLBody:    body,
	}

	err := mail.Send(ctx, msg)

	return err
}
