package usuarios

import (
	"carnegestao/pkg/utils"
	"database/sql"
	"errors"
	"fmt"
)

// db vai ser injetado pelo main.go depois
var db *sql.DB

// Inicializa o repositório com o banco de dados
func InicializarRepositorio(database *sql.DB) {
	db = database
}

// Função para criar usuário
func CriarUsuario(nome, email, senha, tipo string) error {
	// Verifica se o email já existe
	var existe bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM usuarios WHERE email = $1)", email).Scan(&existe)
	if err != nil {
		return fmt.Errorf("erro ao verificar email: %v", err)
	}

	if existe {
		return errors.New("email já cadastrado")
	}

	// Criptografa a senha
	senhaHash, err := utils.CriptografarSenha(senha)
	if err != nil {
		return fmt.Errorf("erro ao criptografar senha: %v", err)
	}

	// Insere no banco
	_, err = db.Exec(`
		INSERT INTO usuarios (nome, email, senha_hash, tipo)
		VALUES ($1, $2, $3, $4)
	`, nome, email, senhaHash, tipo)

	if err != nil {
		return fmt.Errorf("erro ao inserir usuário: %v", err)
	}

	return nil
}
