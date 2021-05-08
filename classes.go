package main

type Cliente struct { // Estrutura básica de um Classe
	nome string
}

type Produto struct {
	nome       string
	quantidade int
	preco      float32
}

type Compra struct {
	produto    Produto
	quantidade int
	valor      float32
}

type Carrinho struct {
	cliente Cliente
	compras []Compra // Array do tipo slice, quando não se declara o tamanho do Array
	valor   float32
}

type Estoque struct {
	produtos []Produto
}
