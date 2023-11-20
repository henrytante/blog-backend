package user

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/repostas"
	"net/http"
	"strings"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nomenick := strings.ToLower(r.URL.Query().Get("usuarios"))
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, err := repositorio.Buscar(nomenick)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, usuarios)
}
