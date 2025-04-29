package clientes

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// Inicializa o repositório com o banco de dados
func InicializarRepositorio(database *sql.DB) {
	db = database
}

// Função para salvar cliente no banco
func SalvarCliente(nomeProprietario, nomeComercial, telefone string, criadoPor int) error {
	_, err := db.Exec(`
		INSERT INTO clientes (nome_proprietario, nome_comercial, telefone, criado_por)
		VALUES ($1, $2, $3, $4)
	`, nomeProprietario, nomeComercial, telefone, criadoPor)

	if err != nil {
		return fmt.Errorf("erro ao salvar cliente: %v", err)
	}

	return nil
}
