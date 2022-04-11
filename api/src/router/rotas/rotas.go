package rotas

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

//Rota reprenta a estrura das rotas da API
type Rota struct {
	URI         string
	Method      string
	Function    func(w http.ResponseWriter, r *http.Request)
	RequireAuth bool
}

//Configure usa o slice da struct de rotas para criar as rotas no padrão do mux
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	// cria o handleFunc("/rota", function).Methods(method) baseado na struct
	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {

			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
