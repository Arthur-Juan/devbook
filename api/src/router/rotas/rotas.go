package rotas

import (
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

	// cria o handleFunc("/rota", function).Methods(method) baseado na struct
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
