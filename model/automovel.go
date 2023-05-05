package model

type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
}

type Moto struct {
	Automovel   // Herança
	Cilindradas int
}

type Carro struct {
	Automovel            // Herança
	QuantidadeDePortas   int
	Potencia             int
	PossuiArCondicionado bool
}
