package vendas

import (
	"carnegestao/internal/auth"
	"encoding/json"
	"net/http"
	"time"
)

type FormaPagamento struct {
	DataPagamento  string  `json:"data_pagamento"`
	Valor          float64 `json:"valor"`
	FormaPagamento string  `json:"forma_pagamento"`
	Observacao     string  `json:"observacao,omitempty"`
}

type VendaItem struct {
	TipoCarne     string  `json:"tipo_carne"`
	Animal        string  `json:"animal"`
	PesoKg        float64 `json:"peso_kg"`
	Quantidade    int     `json:"quantidade"`
	PrecoUnitario float64 `json:"preco_unitario"`
	Lado          string  `json:"lado,omitempty"`
	NumeroAnimal  int     `json:"numero_animal,omitempty"`
}

type VendaRequest struct {
	NumeroNota      int              `json:"numero_nota"`
	Data            string           `json:"data"`
	ClienteID       int              `json:"cliente_id"`
	Desconto        float64          `json:"desconto"`
	Acrescimo       float64          `json:"acrescimo"`
	StatusPagamento string           `json:"status_pagamento"`
	DataQuitacao    string           `json:"data_quitacao,omitempty"`
	Itens           []VendaItem      `json:"itens"`
	Pagamentos      []FormaPagamento `json:"pagamentos"`
}

func CadastrarVendaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req VendaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao ler requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.ClienteID == 0 || req.NumeroNota == 0 {
		http.Error(w, "Cliente e número da nota são obrigatórios", http.StatusBadRequest)
		return
	}

	// Validar ou ajustar a data
	if req.Data == "" {
		req.Data = time.Now().Format("2006-01-02")
	}

	usuarioLogado := r.Context().Value(auth.UsuarioLogadoKey).(auth.UsuarioLogado)

	err = SalvarVenda(req, usuarioLogado.ID)
	if err != nil {
		http.Error(w, "Erro ao salvar venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Venda registrada com sucesso"))
}
