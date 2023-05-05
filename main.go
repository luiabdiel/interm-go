package main

import (
	"fmt"
	"intermgo/model"
	"time"
)

func main() {
	endereco := model.Endereco{
		Rua:    "Rua X",
		Numero: 111,
		Cidade: "Texas",
	}

	pessoa := model.Pessoa{
		Nome:             "Joel",
		Endereco:         endereco,
		DataDeNascimento: time.Date(1981, 9, 24, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)
	endereco.Cidade = "Texas City"
	fmt.Println(endereco.Cidade)
	pessoa.CalculaIdade()
	fmt.Println(pessoa.Idade)
}
