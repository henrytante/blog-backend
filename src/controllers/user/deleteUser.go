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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametro["usuarioID"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	usuarioIDToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if usuarioID != usuarioIDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel deletar um user que não seja o seu"))
		return
	}
	defer db.Close()
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if err = repositorio.Deletar(usuarioID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
