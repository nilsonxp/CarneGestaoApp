package auth

import (
	"carnegestao/pkg/utils"
	"net/http"
	"strings"
)

// Middleware que protege rotas privadas
func AutenticarMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Pega o token do header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		// Espera algo tipo "Bearer eyJhbGciOiJI..."
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato do token inválido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Valida o token
		_, err := utils.ValidarTokenJWT(tokenString)
		if err != nil {
			http.Error(w, "Token inválido: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Token ok → chama o próximo handler
		next.ServeHTTP(w, r)
	}
}
