package utils

import "golang.org/x/crypto/bcrypt"

// CriptografarSenha criptografa uma senha usando bcrypt
func CriptografarSenha(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
