package clientes

import (
	"carnegestao/internal/auth"
	"encoding/json"
	"net/http"
)

type ClienteRequest struct {
	NomeProprietario           string `json:"nome_proprietario"`
	NomeComercial  string `json:"nome_comercial"`
	Telefone       string `json:"telefone"`
}

// Handler para criar novo cliente
func CadastrarClienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req ClienteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	// Validação básica
	if req.NomeProprietario == "" {
		http.Error(w, "Nome do proprietário é obrigatório", http.StatusBadRequest)
		return
	}

	usuarioLogado := r.Context().Value(auth.UsuarioLogadoKey).(auth.UsuarioLogado)

	err = SalvarCliente(req.NomeProprietario, req.NomeComercial, req.Telefone, usuarioLogado.ID)
	if err != nil {
		http.Error(w, "Erro ao salvar cliente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cliente criado com sucesso"))
}
