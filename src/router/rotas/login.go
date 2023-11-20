package rotas

import (
	"api/src/controllers/login"
	"net/http"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             login.Login,
	RequerAutenticacao: false,
}
