package controllers

import "net/http"

// GetBad for http GET
func GetBad() func(w http.ResponseWriter, r *http.Request) {
	return notImplemented
}

// CreateBad for http POST
func CreateBad() func(w http.ResponseWriter, r *http.Request) {
	return notImplemented
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
