package util

import (
	"log"
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrHttp(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
