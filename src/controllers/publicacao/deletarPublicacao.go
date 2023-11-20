package publicacao

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/repositorios"
	respostas "api/src/repostas"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
	}
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parametros["publicacaoID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositorios.NoveRepositorioPublicacoes(db)
	publicacaoSalvaBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if publicacaoSalvaBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Esta publicação não pertence a voce"))
		return
	}
	if err = repositorio.Deletar(publicacaoID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
