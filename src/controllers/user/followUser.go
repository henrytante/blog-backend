package user

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

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Voce não pode seguir a si mesmo"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if err = repositorio.Seguir(usuarioID, seguidorID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
