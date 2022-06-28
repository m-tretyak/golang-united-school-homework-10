package routes

import (
	"github.com/GolangUnited/helloweb/api/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupBadRoute(router *mux.Router) {
	router.
		HandleFunc("/bad", controllers.GetBad()).
		Methods(http.MethodGet)
	router.
		HandleFunc("/bad", controllers.CreateBad()).
		Methods(http.MethodPost)
}
