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

	automovel := model.Automovel{
		Ano:    2023,
		Placa:  "XPTO",
		Modelo: "CG",
	}

	moto := model.Moto{
		Automovel:   automovel,
		Cilindradas: 125,
	}

	fmt.Println(moto)
	fmt.Println(moto.Modelo)
}
