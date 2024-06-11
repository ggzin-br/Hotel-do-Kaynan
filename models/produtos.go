package models

import (
	"log"
	"models/database"
)

// Declarações
type Cliente struct {
	Id int
	Nome string
	CheckIn string
	CheckOut string
	Quarto int
	Proprietario string
}
// Algoritmo

// Login 
/**/
func Login(email string,senha string) bool {

	// Conectar com banco
	db := database.ConectaComBancoDeDados()

	// -> Seleciona todos os usuários
	todos_user, err := db.Query("SELECT * FROM usuarios ORDER BY id ASC")
	if err != nil { // -> Verificação de erros
		panic(err.Error())
	}

		// Checagem dos usuários
		for todos_user.Next() { // -> Enquanto der true o loop roda
		// true = se tiver algum dado nesta linha

		// Declarações
		var temp_email, temp_senha string
		var id int

		// Copiar o valor da linha numa variável temporária
		err = todos_user.Scan(&temp_email,&temp_senha,&id)
			if err != nil {
				log.Println(err.Error())
			}

			if email == temp_email && senha == temp_senha { // -> Se a linha achada foi a mesma do input
				defer db.Close()
				return true
			}

		}
	defer db.Close()
	return false
}
/**/

// Registro
/**/
func Registro(email string,senha string) {

	// Conectar com banco
	db := database.ConectaComBancoDeDados()

	// Preparo da Query
	reg_user, err := db.Prepare("INSERT INTO usuarios(email,senha) VALUES (?,?)")
		if err != nil { // -> Verificação de erros
			panic(err.Error())
		}

		// Executar a query
		reg_user.Exec(email,senha)

	defer db.Close()
}
/**/

// Buscar clientes
/**/
func BuscarClientes(proprietario string) []Cliente { // -> public

	db := database.ConectaComBancoDeDados()

	selectDeTodosOsClientes, err := db.Query("SELECT * FROM clientes WHERE proprietario = '"+proprietario+"' ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	c := Cliente{}
	clientes_slice := []Cliente{}

	for selectDeTodosOsClientes.Next() {
		var id, quarto int
		var nome, checkIn, checkOut, proprietario string

		err = selectDeTodosOsClientes.Scan(&nome, &checkIn, &checkOut, &quarto, &proprietario, &id)
		if err != nil {
			panic(err.Error())
		}

		c.Id = id
		c.Nome = nome
		c.CheckIn = checkIn
		c.CheckOut = checkOut
		c.Quarto = quarto

		clientes_slice = append(clientes_slice, c)
	}

	defer db.Close()
	return clientes_slice
}
/**/

// Reservar
/**/
func Reservar(nome, checkIn string, checkOut string, quarto int, proprietario string) {
	db := database.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO clientes(nome,checkIn,checkOut,quarto,proprietario) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, checkIn, checkOut, quarto, proprietario)
	defer db.Close()
}
/**/

// CheckOut
/**/
func CheckOut(id string) {
	db := database.ConectaComBancoDeDados()

	delete, err := db.Prepare("DELETE FROM clientes WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}
/**/

// Buscar Reservas
/**/
func BuscarReservas(id string) Cliente {
	db := database.ConectaComBancoDeDados()

	clientDB, err := db.Query("SELECT * FROM clientes WHERE id = "+id)

	if err != nil {
		panic(err.Error())
	}

	reservaUpdate := Cliente{}

	for clientDB.Next() {
		var id, quarto int
		var nome, checkIn, checkOut, proprietario string

		err = clientDB.Scan(&nome, &checkIn, &checkOut, &quarto, &proprietario, &id)

		if err != nil {
			panic(err.Error())
		}

		reservaUpdate.Id = id
		reservaUpdate.Nome = nome
		reservaUpdate.CheckIn = checkIn
		reservaUpdate.CheckOut = checkOut
		reservaUpdate.Quarto = quarto
	}

	defer db.Close()

	return reservaUpdate
}
/**/

// Mod Cliente
/**/
func AtualizacaoCliente(id int, nome, checkIn string, checkOut string, quarto int) {
	db := database.ConectaComBancoDeDados()

	atualizarCliente, err := db.Prepare("UPDATE clientes SET nome = ?, checkIn = ?, checkOut = ?, quarto = ? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	atualizarCliente.Exec(nome, checkIn, checkOut, quarto, id)

	defer db.Close()
}
/**/
