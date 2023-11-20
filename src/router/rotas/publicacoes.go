package rotas

import (
	"api/src/controllers/publicacao"
	"net/http"
)

var rotaPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             publicacao.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             publicacao.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{publicacaoID}",
		Metodo:             http.MethodGet,
		Funcao:             publicacao.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{publicacaoID}",
		Metodo:             http.MethodPut,
		Funcao:             publicacao.EditarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoID}",
		Metodo:             http.MethodDelete,
		Funcao:             publicacao.DeletarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{usuarioID}/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             publicacao.UserPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{publicacaoID}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             publicacao.Curtir,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{publicacaoID}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             publicacao.Descurtir,
		RequerAutenticacao: true,
	},
}
