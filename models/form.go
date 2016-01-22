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
		this.Errors["Email"] = "Veuillez entrer une adresse email valide."
	}

	if strings.TrimSpace(this.Name) == "" {
		this.Errors["Name"] = "Veuillez entrer votre nom dans les champs vierges."
	}

	if strings.TrimSpace(this.Phone) == "" {
		this.Errors["Phone"] = "Veuillez entrer votre numéro de téléphone."
	}

	if strings.TrimSpace(this.Message) == "" {
		this.Errors["Message"] = "Veuillez entrer un message."
	}

	validated := len(this.Errors) == 0
	if validated {
		this.Thanks = "Merci pour votre demande! Je vais vous contacter bientôt."
	}

	return validated
}
