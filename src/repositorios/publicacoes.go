package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NoveRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}
	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoIDInserido), nil
}
func (repositorio Publicacoes) BuscarPorID(ID uint64) (modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query("SELECT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id = p.autor_id WHERE p.id = ?", ID)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao
	if linhas.Next() {
		if err = linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNICK); err != nil {
			return modelos.Publicacao{}, err
		}
	}
	return publicacao, nil
}
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p 
	inner join usuarios u on u.id = p.autor_id 
	inner join seguidores s on p.autor_id = s.usuario_id 
	where u.id = ? or s.seguidor_id = ?
	order by 1 desc`,
		usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNICK,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, err := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}
	return nil
}
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}
	return nil
}
func (repositorio Publicacoes) BuscarPorUsuario(ID uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p
		join usuarios u on u.id = p.autor_id
		where p.autor_id = ?`,
		ID,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao

		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNICK,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Curtir(ID uint64) error {
	statement, err := repositorio.db.Prepare("UPDATE publicacoes set curtidas = curtidas + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (repositorio Publicacoes) Descurtir(ID uint64) error {
	statement, err := repositorio.db.Prepare(`
		update publicacoes set curtidas = 
		CASE 
			WHEN curtidas > 0 THEN curtidas - 1
			ELSE 0 
		END
		where id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}
