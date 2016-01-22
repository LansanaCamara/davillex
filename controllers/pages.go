package controllers

import (
	"github.com/lansanacamara/davillex/models"
	"github.com/lansanacamara/davillex/util"
	"fmt"
	"net/http"
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
	util.HTML(w, "index", nil)
}

func doPost(w http.ResponseWriter, r *http.Request) {
	form := &models.Form{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Phone:   r.FormValue("phone"),
		Message: r.FormValue("message"),
	}

	if form.Validate() {
		from := form.Email
		message := fmt.Sprintf(util.CONTACT_EMAIL, form.Name, form.Email, form.Email, form.Phone, form.Phone, form.Message)

		err := util.SendMail(from, message)
		util.CheckErrHttp(w, err)
	}

	util.HTML(w, "index", form)
}
