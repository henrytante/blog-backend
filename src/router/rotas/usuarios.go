package rotas

import (
	"api/src/controllers/user"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             user.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             user.SearchUsers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}",
		Metodo:             http.MethodGet,
		Funcao:             user.SearchUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}",
		Metodo:             http.MethodPut,
		Funcao:             user.UpdateUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}",
		Metodo:             http.MethodDelete,
		Funcao:             user.DeleteUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}/follow",
		Metodo:             http.MethodPost,
		Funcao:             user.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}/unfollow",
		Metodo:             http.MethodPost,
		Funcao:             user.Unfollow,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}/followers",
		Metodo:             http.MethodGet,
		Funcao:             user.Followers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}/following",
		Metodo:             http.MethodGet,
		Funcao:             user.Following,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}/passupdate",
		Metodo:             http.MethodPost,
		Funcao:             user.PassUpdate,
		RequerAutenticacao: true,
	},
}
