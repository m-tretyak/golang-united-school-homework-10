package routes

import (
	"fmt"
	"github.com/GolangUnited/helloweb/api/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupNameRoute(router *mux.Router) {
	router.
		HandleFunc(fmt.Sprintf("/name/{%s}", controllers.ParamsUriPartName), controllers.GetName()).
		Methods(http.MethodGet)
}
