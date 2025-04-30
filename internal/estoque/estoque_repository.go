package estoque

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InicializarRepositorio(database *sql.DB) {
	db = database
}

func RegistrarEntradaEstoque(req EntradaEstoqueRequest, criadoPor int) error {
	_, err := db.Exec(`
		INSERT INTO estoque (
			data, quantidade_bois, quantidade_vacas,
			peso_total_bois, peso_total_vacas,
			visceras_recebidas, visceras_condenadas,
			fornecedor, criado_por
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`,
		req.Data,
		req.QuantidadeBois,
		req.QuantidadeVacas,
		req.PesoBois,
		req.PesoVacas,
		req.ViscerasRecebidas,
		req.ViscerasCondenadas,
		req.Fornecedor,
		criadoPor,
	)
	if err != nil {
		return fmt.Errorf("erro ao inserir estoque: %w", err)
	}
	return nil
}

type EntradaEstoque struct {
	ID                 int     `json:"id"`
	Data               string  `json:"data"`
	QuantidadeBois     int     `json:"quantidade_bois"`
	QuantidadeVacas    int     `json:"quantidade_vacas"`
	PesoBois           float64 `json:"peso_total_bois"`
	PesoVacas          float64 `json:"peso_total_vacas"`
	ViscerasRecebidas  int     `json:"visceras_recebidas"`
	ViscerasCondenadas int     `json:"visceras_condenadas"`
	Fornecedor         string  `json:"fornecedor"`
	CriadoEm           string  `json:"criado_em"`
	CriadoPor          int     `json:"criado_por"`
}

func ListarEntradasEstoque() ([]EntradaEstoque, error) {
	rows, err := db.Query(`
		SELECT id, data, quantidade_bois, quantidade_vacas,
		       peso_total_bois, peso_total_vacas, visceras_recebidas,
		       visceras_condenadas, fornecedor, criado_em, criado_por
		FROM estoque
		ORDER BY data DESC, id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []EntradaEstoque
	for rows.Next() {
		var e EntradaEstoque
		err := rows.Scan(
			&e.ID, &e.Data, &e.QuantidadeBois, &e.QuantidadeVacas,
			&e.PesoBois, &e.PesoVacas, &e.ViscerasRecebidas,
			&e.ViscerasCondenadas, &e.Fornecedor, &e.CriadoEm, &e.CriadoPor,
		)
		if err != nil {
			return nil, err
		}
		lista = append(lista, e)
	}

	return lista, nil
}
