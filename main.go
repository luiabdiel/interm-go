package main

import (
	"fmt"
	"intermgo/model"
)

func main() {
	endereco := model.Endereco{
		Rua:    "Rua X",
		Numero: 111,
		Cidade: "Tilapia",
	}

	fmt.Println(endereco)
	endereco.Cidade = "Tilapinha do Norte"
	fmt.Println(endereco.Cidade)
}
