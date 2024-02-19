package entities

type Cliente struct { // Estrutura básica de um Classe
	Nome string
}

type Produto struct {
	Nome       string
	Quantidade int
	Preco      float32
}

type Compra struct {
	Produto    Produto
	Quantidade int
	Valor      float32
}

type Carrinho struct {
	Cliente Cliente
	Compras []Compra // Array do tipo slice, quando não se declara o tamanho do Array
	Valor   float32
}

type Estoque struct {
	Produtos []Produto
}
