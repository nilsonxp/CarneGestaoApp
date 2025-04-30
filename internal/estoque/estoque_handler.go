package estoque

import (
	"carnegestao/internal/auth"
	"encoding/json"
	"net/http"
)

type EntradaEstoqueRequest struct {
	Data               string  `json:"data"`
	QuantidadeBois     int     `json:"quantidade_bois"`
	QuantidadeVacas    int     `json:"quantidade_vacas"`
	PesoBois           float64 `json:"peso_total_bois"`
	PesoVacas          float64 `json:"peso_total_vacas"`
	ViscerasRecebidas  int     `json:"visceras_recebidas"`
	ViscerasCondenadas int     `json:"visceras_condenadas"`
	Fornecedor         string  `json:"fornecedor"`
}

func CadastrarEstoqueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req EntradaEstoqueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao ler requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validação simples
	if req.Data == "" || (req.QuantidadeBois == 0 && req.QuantidadeVacas == 0) {
		http.Error(w, "Data e pelo menos uma quantidade devem ser informadas", http.StatusBadRequest)
		return
	}

	usuarioLogado := r.Context().Value(auth.UsuarioLogadoKey).(auth.UsuarioLogado)

	req.ViscerasRecebidas = req.QuantidadeBois + req.QuantidadeVacas

	err = RegistrarEntradaEstoque(req, usuarioLogado.ID)
	if err != nil {
		http.Error(w, "Erro ao salvar estoque: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Estoque cadastrado com sucesso"))
}

func ListarEstoqueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	estoque, err := ListarEntradasEstoque()
	if err != nil {
		http.Error(w, "Erro ao buscar estoque: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estoque)
}
