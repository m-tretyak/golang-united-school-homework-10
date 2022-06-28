package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const ParamsUriPartName = "PARAMS"

// GetName for http GET
func GetName() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := mux.Vars(r)[ParamsUriPartName]

		_, err := fmt.Fprintf(w, "Hello, %s!", value)

		if err != nil {
			log.Println(err)
		}
	}
}

// CreateName for http POST
func CreateName() func(w http.ResponseWriter, r *http.Request) {
	return notImplemented
}
