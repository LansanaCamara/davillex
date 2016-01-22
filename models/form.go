package models

import (
	"regexp"
	"strings"
)

type Form struct {
	Name    string
	Email   string
	Phone   string
	Message string
	Thanks  string
	Errors  map[string]string
}

func (this *Form) Validate() bool {
	this.Errors = make(map[string]string)

	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(this.Email))
	if matched == false {
		this.Errors["Email"] = "Please enter a valid email address."
	}

	if strings.TrimSpace(this.Name) == "" {
		this.Errors["Name"] = "Please enter your name."
	}

	if strings.TrimSpace(this.Phone) == "" {
		this.Errors["Phone"] = "Please enter your phone number."
	}

	if strings.TrimSpace(this.Message) == "" {
		this.Errors["Message"] = "Please enter a message."
	}

	validated := len(this.Errors) == 0
	if validated {
		this.Thanks = "Thanks for your request! Your email has been sent."
	}

	return validated
}
