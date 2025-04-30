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

type Cliente struct {
	ID             int    `json:"id"`
	NomeProprietario string `json:"nome_proprietario"`
	NomeComercial   string `json:"nome_comercial"`
	Telefone        string `json:"telefone"`
	CriadoEm         string `json:"criado_em"`
	CriadoPor         int    `json:"criado_por"`
}

// Buscar todos os clientes do banco
func ListarClientes() ([]Cliente, error) {
	rows, err := db.Query(`
		SELECT id, nome_proprietario, nome_comercial, telefone, criado_em, criado_por
		FROM clientes
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []Cliente
	for rows.Next() {
		var c Cliente
		err := rows.Scan(&c.ID, &c.NomeProprietario, &c.NomeComercial, &c.Telefone, &c.CriadoEm, &c.CriadoPor)
		if err != nil {
			return nil, err
		}
		lista = append(lista, c)
	}

	return lista, nil
}