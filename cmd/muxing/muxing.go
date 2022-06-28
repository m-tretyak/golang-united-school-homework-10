package main

import (
	"fmt"
	"github.com/GolangUnited/helloweb/api/helper"
	"log"
	"net/http"
	"os"
	"strconv"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := helper.NewRouter()

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	//if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), helper.LowerCaseURI(router)); err != nil {
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}

	Start(host, port)
}
