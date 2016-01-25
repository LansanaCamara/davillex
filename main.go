package main

import (
	"net/http"

	"github.com/lansanacamara/davillex/controllers"
	"github.com/lansanacamara/davillex/models"
)

func init() {
	cssHandler := http.FileServer(http.Dir("./public/css/"))
	jsHandler := http.FileServer(http.Dir("./public/js/"))
	imgHandler := http.FileServer(http.Dir("./public/img/"))
	fontHandler := http.FileServer(http.Dir("./public/font-awesome/"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandler))
	http.Handle("/img/", http.StripPrefix("/img/", imgHandler))
	http.Handle("/font-awesome/", http.StripPrefix("/font-awesome/", fontHandler))

	http.Handle("/", models.Handle(controllers.Index))
}