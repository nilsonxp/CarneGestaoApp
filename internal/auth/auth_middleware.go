package auth

import (
	"carnegestao/pkg/utils"
	"net/http"
	"strings"
	"context"
)

type contextKey string

const UsuarioLogadoKey contextKey = "usuarioLogado"

type UsuarioLogado struct {
	ID int
	Tipo string
}

// Middleware que protege rotas privadas
func AutenticarMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato do token inválido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := utils.ValidarTokenJWT(tokenString)
		if err != nil {
			http.Error(w, "Token inválido: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Salvar dados do usuário logado no contexto
		ctx := context.WithValue(r.Context(), UsuarioLogadoKey, UsuarioLogado{
			ID:   claims.IDUsuario,
			Tipo: claims.Tipo,
		})

		// Chama o próximo handler com o contexto atualizado
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Verifica se o usuário logado é admin
func SomenteAdmin(next http.HandlerFunc) http.HandlerFunc {
	return AutenticarMiddleware(func(w http.ResponseWriter, r *http.Request) {
		usuarioLogado := r.Context().Value(UsuarioLogadoKey).(UsuarioLogado)

		if usuarioLogado.Tipo != "admin" {
			http.Error(w, "Acesso restrito a administradores", http.StatusForbidden)
			return
		}

		next(w, r)
	})
}

// Verifica se o usuário logado é admin OU funcionário
func SomenteAdminOuFuncionario(next http.HandlerFunc) http.HandlerFunc {
	return AutenticarMiddleware(func(w http.ResponseWriter, r *http.Request) {
		usuarioLogado := r.Context().Value(UsuarioLogadoKey).(UsuarioLogado)

		if usuarioLogado.Tipo != "admin" && usuarioLogado.Tipo != "funcionario" {
			http.Error(w, "Acesso restrito a administradores e funcionários", http.StatusForbidden)
			return
		}

		next(w, r)
	})
}
