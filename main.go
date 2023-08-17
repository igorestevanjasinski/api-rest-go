package main

import (
	"fmt"

	"github.com/igorestevanjasinski/api-rest-go/database"
	"github.com/igorestevanjasinski/api-rest-go/models"
	"github.com/igorestevanjasinski/api-rest-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
