package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
func (repositorio Usuarios) BuscarID(ID uint64) (modelos.Usuario, error) {
	linha, err := repositorio.db.Query("SELECT id, nome, nick, email, criadoEm from usuarios where id = ?", ID)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()
	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}
	return nil
}
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, err := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()
	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}
	return nil
}
func (repositorio Usuarios) Unfollow(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}
	return nil
}
func (repositorio Usuarios) Followers(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, err := repositorio.db.Query("SELECT u.id, u.nome, u.nick,u.email,u.criadoEm from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?", usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
func (repositorio Usuarios) Following(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, err := repositorio.db.Query("select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?", usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
func (repositorio Usuarios) BuscarPass(usuarioID uint64) (string, error) {
	linha, err := repositorio.db.Query("SELECT senha FROM usuarios WHERE id = ?", usuarioID)
	if err != nil {
		return "", err
	}
	defer linha.Close()
	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Senha); err != nil {
			return "", err
		}
	}
	return usuario.Senha, nil
}
func (repositorio Usuarios) UpdatePass(ID uint64, NewPass string) error {
	statement, err := repositorio.db.Prepare("UPDATE  usuarios SET senha = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(NewPass, ID); err != nil {
		return err
	}
	return nil
}
