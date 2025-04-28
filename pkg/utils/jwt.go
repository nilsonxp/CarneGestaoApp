package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

var chaveSecreta = []byte("chave-secreta-muito-forte") // depois mover para .env

type Claims struct {
	IDUsuario int    `json:"id_usuario"`
	Tipo      string `json:"tipo"`
	jwt.RegisteredClaims
}

// Gera um token JWT para o usuário
func GerarTokenJWT(idUsuario int, tipo string) (string, error) {
	expiracao := time.Now().Add(24 * time.Hour) // token válido por 24h

	claims := Claims{
		IDUsuario: idUsuario,
		Tipo:      tipo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiracao),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(chaveSecreta)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
