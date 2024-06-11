package login

import (
	"net/http"
)

var Auth_login = false

// Verificador de login
func VerLogin(rota http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !Auth_login {
			http.Error(w, "Usuário não autenticado", http.StatusForbidden)
			return
		} else {
			rota.ServeHTTP(w, r)
		}
	})
}

/*
Deixarei este arquivo aqui por questões de arquivamento

Uma vez que o login é verdadeiro, a rota se abre e não se fecha mais
Não pretendo avançar em um sistema muito complexo
*/