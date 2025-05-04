package vendas

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InicializarRepositorio(database *sql.DB) {
	db = database
}

func SalvarVenda(req VendaRequest, criadoPor int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}
	defer tx.Rollback()

	valorTotal := 0.0
	for _, item := range req.Itens {
		valorTotal += item.PrecoUnitario * float64(item.Quantidade)
	}
	valorTotal -= req.Desconto
	valorTotal += req.Acrescimo

	// Trata data_quitacao se vier vazia
	var dataQuitacao interface{} = nil
	if req.DataQuitacao != "" {
		dataQuitacao = req.DataQuitacao
	}

	var vendaID int
	err = tx.QueryRow(`
		INSERT INTO vendas (
			cliente_id, data, numero_nota,
			total_final, desconto, acrescimo,
			status_pagamento, data_quitacao,
			criado_por
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`,
		req.ClienteID, req.Data, req.NumeroNota,
		valorTotal, req.Desconto, req.Acrescimo,
		req.StatusPagamento, dataQuitacao,
		criadoPor,
	).Scan(&vendaID)
	if err != nil {
		return fmt.Errorf("erro ao inserir venda: %w", err)
	}

	for _, item := range req.Itens {
		precoTotal := item.PrecoUnitario * float64(item.Quantidade)
		_, err = tx.Exec(`
			INSERT INTO venda_itens (
				venda_id, produto_id, tipo_animal,
				peso_kg, quantidade, preco_unitario,
				preco_total, lado, numero_animal
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`,
			vendaID, item.ProdutoID, item.TipoAnimal,
			item.PesoKg, item.Quantidade, item.PrecoUnitario,
			precoTotal, item.Lado, item.NumeroAnimal,
		)
		if err != nil {
			return fmt.Errorf("erro ao inserir item da venda: %w", err)
		}
	}

	for _, pagamento := range req.Pagamentos {
		_, err = tx.Exec(`
			INSERT INTO venda_pagamentos (
				venda_id, data_pagamento,
				valor_pagamento, forma_pagamento,
				observacao, criado_por
			)
			VALUES ($1, $2, $3, $4, $5, $6)
		`,
			vendaID, pagamento.DataPagamento,
			pagamento.Valor, pagamento.FormaPagamento,
			pagamento.Observacao, criadoPor,
		)
		if err != nil {
			return fmt.Errorf("erro ao inserir pagamento: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("erro ao finalizar transação: %w", err)
	}

	return nil
}

func SomaPagamentos(formas []FormaPagamento) float64 {
	total := 0.0
	for _, forma := range formas {
		total += forma.Valor
	}
	return total
}

func BuscarPrecosDoDia(data string) (map[string]float64, error) {
	return map[string]float64{
		"casada_boi":   10.0,
		"casada_vaca":  9.0,
		"dianteiro":    6.0,
		"traseiro":     8.0,
		"viscera":      5.0,
		"panelada":     7.0,
	}, nil
}
