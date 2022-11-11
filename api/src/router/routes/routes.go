package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Todas rotas da API
type Route struct {
	Uri         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := UserRoutes

	for _, routes := range routes {
		r.HandleFunc(routes.Uri, routes.Function).Methods(routes.Method)
	}
	return r
}
