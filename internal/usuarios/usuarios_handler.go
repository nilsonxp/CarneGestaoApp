package usuarios

import (
	"carnegestao/internal/auth"
	"encoding/json"
	"net/http"
)

type UsuarioRequest struct {
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Tipo     string `json:"tipo"`
}

// Handler para criar novo usuário
func CriarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req UsuarioRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	// Validação básica
	if req.Nome == "" || req.Email == "" || req.Senha == "" || req.Tipo == "" {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	if req.Tipo != "admin" && req.Tipo != "funcionario" && req.Tipo != "cliente" {
		http.Error(w, "Tipo inválido", http.StatusBadRequest)
		return
	}

	usuarioLogado := r.Context().Value(auth.UsuarioLogadoKey).(auth.UsuarioLogado)

	// Chama função do repositório para salvar no banco
	err = CriarUsuario(req.Nome, req.Email, req.Senha, req.Tipo, usuarioLogado.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário criado com sucesso"))
}