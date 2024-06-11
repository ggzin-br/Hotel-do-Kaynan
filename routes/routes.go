package routes

import (
	"models/controllers"
	"net/http"
)

var Mux *http.ServeMux

func Rotas() {
	// Início

	Mux = http.NewServeMux()

	// Pag
	Mux.Handle("/home", http.HandlerFunc(controllers.Index))
	Mux.Handle("/home/criar-reserva", http.HandlerFunc(controllers.CriarReserva))
	Mux.Handle("/home/editar", http.HandlerFunc(controllers.Atualizar))

	// POST
	Mux.Handle("/home/SQL_CriarReserva", http.HandlerFunc(controllers.SQL_CriarReserva))
	Mux.Handle("/home/SQL_Deletar", http.HandlerFunc(controllers.SQL_Deletar))
	Mux.Handle("/home/SQL_Atualizar", http.HandlerFunc(controllers.SQL_Atualizar))

	/**/
	// Login & Registro

	// Pag
	Mux.Handle("/", http.HandlerFunc(controllers.Login))
	// POST
	Mux.Handle("/SQL_Login", http.HandlerFunc(controllers.SQL_Login))

	// Pag
	Mux.Handle("/cria-user", http.HandlerFunc(controllers.CriarUser))
	// POST
	Mux.Handle("/SQL_CriarUser", http.HandlerFunc(controllers.SQL_CriarUser))

	// POST
	Mux.Handle("/SERVER_LogOut", http.HandlerFunc(controllers.LogOut))

}
