package controllers

import (
	"fmt"
	"net/http"

	"github.com/lansanacamara/davillex/config"
	"github.com/lansanacamara/davillex/models"
	"github.com/lansanacamara/davillex/util"
)

func Index(w http.ResponseWriter, r *http.Request) *models.Error {
	if r.Method == "GET" {
		util.RenderView(w, "index", nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		form := &models.Form{
			Name:    util.FetchAndClean(r, "name"),
			Email:   util.FetchAndClean(r, "email"),
			Phone:   util.FetchAndClean(r, "phone"),
			Message: util.FetchAndClean(r, "message"),
		}

		if form.Validate() {
			body := fmt.Sprintf(config.CONTACT_EMAIL, form.Name, form.Email, form.Phone, form.Message)

			if err := util.SendMail(body); err != nil {
				return &models.Error{
					Error: err,
					Code: 500,
					Message: "Votre lettre pas pu être envoyé. Veuillez réessayer plus tard.",
					View: "index",
				}
			} else {
				form.Thanks = config.THANK_YOU
			}
		}

		util.RenderView(w, "index", form)
	}
	return nil
}