package user

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	respostas "api/src/repostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func PassUpdate(w http.ResponseWriter, r *http.Request) {
	usuarioIDToken, err := autenticacao.ExtrairUsuarioID(r)
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
	if usuarioID != usuarioIDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Voce não pode mudar a senha de outro user"))
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var pass modelos.Pass
	if err = json.Unmarshal(reqBody, &pass); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	passBanco, err := repositorio.BuscarPass(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if err = seguranca.VerificarSenha(passBanco, pass.Current); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("As senhas não batem"))
		return
	}
	senhaHash, err := seguranca.Hash(pass.New)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = repositorio.UpdatePass(usuarioID, string(senhaHash)); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
