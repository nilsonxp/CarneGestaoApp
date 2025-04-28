package auth

import (
	// "carnegestao/internal/usuarios"
	"carnegestao/pkg/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

var db *sql.DB

func InicializarAuth(database *sql.DB) {
	db = database
}

type LoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Senha == "" {
		http.Error(w, "Email e senha obrigatórios", http.StatusBadRequest)
		return
	}

	// Buscar usuário no banco
	var id int
	var senhaHash string
	var tipo string
	err = db.QueryRow(`
		SELECT id, senha_hash, tipo FROM usuarios WHERE email = $1
	`, req.Email).Scan(&id, &senhaHash, &tipo)

	if err == sql.ErrNoRows {
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
		return
	}

	// Verificar senha
	if !utils.CompararSenha(req.Senha, senhaHash) {
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
		return
	}

	// Gerar token
	token, err := utils.GerarTokenJWT(id, tipo)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	resp := LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
