package controllers

import (
	"fmt"
	"net/http"

	"github.com/lansanacamara/davillex/config"
	"github.com/lansanacamara/davillex/models"
	"github.com/lansanacamara/davillex/util"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		doGet(w, r)
	case "POST":
		doPost(w, r)
	default:
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}
}

func doGet(w http.ResponseWriter, r *http.Request) {
	util.RenderView(w, "index", nil)
}

func doPost(w http.ResponseWriter, r *http.Request) {
	form := &models.Form{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Phone:   r.FormValue("phone"),
		Message: r.FormValue("message"),
	}

	if form.Validate() {
		body := fmt.Sprintf(config.CONTACT_EMAIL, form.Name, form.Email, form.Phone, form.Message)

		if err := util.SendMail(body, r); err != nil {
			form.HandleErr(config.MAIL_NOT_SENT_ERR)
		} else {
			form.Thanks = config.THANK_YOU
		}
	}

	util.RenderView(w, "index", form)
}
