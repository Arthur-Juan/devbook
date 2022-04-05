package router

import (
	"api/src/router/rotas"
	"github.com/gorilla/mux"
)

//Init vai retornar um router com as rotas configuradas
func Init() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configure(r) //envia um newRouter para a função configure
}
