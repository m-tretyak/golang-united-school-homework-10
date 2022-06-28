package routes

import (
	"github.com/GolangUnited/helloweb/api/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupHeadersRoute(router *mux.Router) {
	router.
		HandleFunc("/headers", controllers.CreateHeaders()).
		Methods(http.MethodPost)
}
