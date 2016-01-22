package util

import (
	"html/template"
	"net/http"
)

func HTML(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles("public/views/" + filename + ".html")
	CheckErrHttp(w, err)

	err = tmpl.Execute(w, data)
	CheckErrHttp(w, err)
}
