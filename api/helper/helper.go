package helper

import (
	"github.com/GolangUnited/helloweb/api/routes"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func NewRouter() (result *mux.Router) {
	result = mux.NewRouter()

	routes.SetupBadRoute(result)
	routes.SetupNameRoute(result)
	routes.SetupDataRoute(result)
	routes.SetupHeadersRoute(result)

	return result
}

func LowerCaseURI(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
