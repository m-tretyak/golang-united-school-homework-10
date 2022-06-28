package routes

import (
	"github.com/GolangUnited/helloweb/api/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupDataRoute(router *mux.Router) {
	router.
		HandleFunc("/data", controllers.CreateData()).
		Methods(http.MethodPost)
}
