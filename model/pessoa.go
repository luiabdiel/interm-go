package model

import "time"

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

// SEMPRE que utilizar métodos para alterar dados de uma struct LEMBRE-SE sempre de dizer que é um ponteiro!!
func (pessoa *Pessoa) CalculaIdade() {
	anoDeNascimento := pessoa.DataDeNascimento.Year()
	anoAtual := time.Now().Year()

	pessoa.Idade = anoAtual - anoDeNascimento
}
