package model

import "time"

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
}

func (pessoa Pessoa) IdadeAtual() int {
	anoDeNascimento := pessoa.DataDeNascimento.Year()
	anoAtual := time.Now().Year()

	return anoAtual - anoDeNascimento
}
