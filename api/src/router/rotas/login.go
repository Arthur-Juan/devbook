package rotas

import (
	"api/src/controllers/loginController"
	"net/http"
)

var loginRoute = Rota{
	URI:         "/login",
	Method:      http.MethodPost,
	Function:    loginController.Login,
	RequireAuth: false,
}
