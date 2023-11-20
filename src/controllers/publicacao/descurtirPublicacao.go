package publicacao

import (
	"api/src/banco"
	"api/src/repositorios"
	respostas "api/src/repostas"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Descurtir(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	pubID, err := strconv.ParseUint(parametro["publicacaoID"], 10, 64)
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
	if err = repositorio.Descurtir(pubID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
