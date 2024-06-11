package controllers

import (
	"html/template"
	"log"
	"models/models"
	"net/http"
	"strconv"
)

// Concatenação
var temp = template.Must(template.ParseGlob("templates/*.html"))

// User id
var User_login string

// Login
/**/
func Login(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "login", nil)
}
func LogOut(w http.ResponseWriter, r *http.Request) {
	User_login = "" // Apaga o registro anterior
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func SQL_Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("campo_email")
		senha := r.FormValue("campo_senha")

		if models.Login(email, senha) {
			User_login = email
			http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		} else {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
	}
}
/**/

// Registro
func CriarUser(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "registro", nil)
}
func SQL_CriarUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("campo_email")
		senha := r.FormValue("campo_senha")

		models.Registro(email, senha)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

/**/

// Início
/**/
func Index(w http.ResponseWriter, r *http.Request) {
	clientes := models.BuscarClientes(User_login)
	temp.ExecuteTemplate(w, "Index", clientes)
}

/**/

// Criação de dados
/**/
func CriarReserva(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "criar-reserva", nil)
}
func SQL_CriarReserva(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		checkIn := r.FormValue("checkin")
		checkOut := r.FormValue("checkout")
		quarto := r.FormValue("quarto")

		quartoInt, err := strconv.Atoi(quarto)
		if err != nil {
			log.Println(err)
		}

		models.Reservar(nome, checkIn, checkOut, quartoInt, User_login)
	}

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

/**/

// Atualização
/**/
func Atualizar(w http.ResponseWriter, r *http.Request) {
	idCliente := r.URL.Query().Get("id")
	cliente := models.BuscarReservas(idCliente)

	temp.ExecuteTemplate(w, "Edit", cliente)
}
func SQL_Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		checkIn := r.FormValue("checkin")
		checkOut := r.FormValue("checkout")
		quarto := r.FormValue("quarto")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}

		quartoInt, err := strconv.Atoi(quarto)
		if err != nil {
			log.Println(err)
		}

		models.AtualizacaoCliente(idConv, nome, checkIn, checkOut, quartoInt)

	}

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

/**/

// Deletar
func SQL_Deletar(w http.ResponseWriter, r *http.Request) {
	idCliente := r.URL.Query().Get("id")
	models.CheckOut(idCliente)
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}
