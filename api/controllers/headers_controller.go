package controllers

import (
	"log"
	"net/http"
	"strconv"
)

// GetHeaders for http GET
func GetHeaders() func(w http.ResponseWriter, r *http.Request) {
	return notImplemented
}

const (
	varAName      = "a"
	varBName      = "b"
	varResultName = varAName + "+" + varBName
)

// CreateHeaders for http POST
func CreateHeaders() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		a, err := strconv.Atoi(r.Header.Get(varAName))

		if err != nil {
			log.Println(err)
		}

		b, err := strconv.Atoi(r.Header.Get(varBName))

		if err != nil {
			log.Println(err)
		}

		w.Header().Set(varResultName, strconv.Itoa(a+b))
	}
}
