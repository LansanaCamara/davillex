package util

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"os"

	"github.com/go-gomail/gomail"
	_ "github.com/joho/godotenv/autoload"
)

func FetchAndClean(r *http.Request, formName string) string {
	return html.EscapeString(r.Form.Get(formName))
}

func MatchRegexp(expression string, str string) bool {
	re := regexp.MustCompile(expression)
	return re.Match([]byte(str))
}

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

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

type EmailDetails struct {
	EmailUsername string
	EmailPassword string
	EmailTwo      string
}

func SendMail(body string) error {
	e := &EmailDetails{
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("EMAIL_TWO"),
	}

	m := gomail.NewMessage()

	m.SetHeader("From", e.EmailUsername)
	m.SetHeader("To", e.EmailUsername, e.EmailTwo)
	m.SetHeader("Subject", "New Message from Davillex.com!")
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, e.EmailUsername, e.EmailPassword)

	err := d.DialAndSend(m)

	return err
}