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
	// Iniciar transação
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}
	defer tx.Rollback() // Rollback em caso de erro

	// 1. Calcular valor total dos itens
	valorTotal := 0.0
	for _, item := range req.Itens {
		valorTotal += item.PrecoUnitario * float64(item.Quantidade)
	}
	valorTotal -= req.Desconto
	valorTotal += req.Acrescimo

	// 2. Inserir venda principal
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
		req.StatusPagamento, req.DataQuitacao,
		criadoPor,
	).Scan(&vendaID)

	if err != nil {
		return fmt.Errorf("erro ao inserir venda: %w", err)
	}

	// 3. Inserir itens da venda
	for _, item := range req.Itens {
		_, err = tx.Exec(`
			INSERT INTO venda_itens (
				venda_id, tipo_carne, animal,
				peso_kg, quantidade, preco_unitario,
				preco_total, lado, numero_animal
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`,
			vendaID, item.TipoCarne, item.Animal,
			item.PesoKg, item.Quantidade, item.PrecoUnitario,
			item.PrecoUnitario*float64(item.Quantidade),
			item.Lado, item.NumeroAnimal,
		)
		if err != nil {
			return fmt.Errorf("erro ao inserir item da venda: %w", err)
		}
	}

	// 4. Inserir pagamentos
	for _, pagamento := range req.Pagamentos {
		_, err = tx.Exec(`
			INSERT INTO venda_pagamentos (
				venda_id, data_pagamento,
				valor_pagamento, forma_pagamento,
				observacao
			)
			VALUES ($1, $2, $3, $4, $5)
		`,
			vendaID, pagamento.DataPagamento,
			pagamento.Valor, pagamento.FormaPagamento,
			pagamento.Observacao,
		)
		if err != nil {
			return fmt.Errorf("erro ao inserir pagamento: %w", err)
		}
	}

	// Commit da transação
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
	// Simulação por enquanto
	return map[string]float64{
		"casada_boi":  10.0,
		"casada_vaca": 9.0,
		"dianteiro":   6.0,
		"traseiro":    8.0,
		"viscera":     5.0,
		"panelada":    7.0,
	}, nil
}
