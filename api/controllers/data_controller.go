package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// GetData for http GET
func GetData() func(w http.ResponseWriter, r *http.Request) {
	return notImplemented
}

// CreateData for http POST
func CreateData() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println(err)
		}

		_, err = fmt.Fprintf(w, "I got message:\n%s", data)

		if err != nil {
			log.Println(err)
		}
	}
}
