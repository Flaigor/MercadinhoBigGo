package main

import (
	"fmt"
	"math"
)

func Saudacoes(nome string) { // declaração base de uma func
	fmt.Println("Olá", nome, ", seja bem vindo!")
}

func MontaMenurPrincipal() {
	fmt.Println("|| 1 - Adicionar produto ao carrinho\n|| 2 - Excluir prduto do carrinho")
	fmt.Println("|| 3 - Consultar Carrinho\n|| 4 - Fechar Pedido\n|| 0 - Fazer Logoff")
}

func CalculaQuadradoECubo(x rune) (int, int) { // É possivel retornar mais de um valor em uma func
	// rune = int32 porém precisa ser convertido, caso não seja int32
	return int((x * x)), int((x * x * x))
}

func CalcularDoisValores(x float32, y float32, op uint8) (resp float32) { // Valores de retorno nomeados não são recomendados para funcs muito grandes
	// Outra forma de fazer switch é colocando "true" n
	switch op { // true
	case 1: // op == 1
		resp = x + y
		return // Como o "resp" já foi declarado como retorno, ele fica implícito sempre que tem um "return"
	case 2: // op == 2
		resp = x - y
		return
	case 3: // op == 3
		resp = x * y
		return
	case 4: // op == 4
		resp = x / y
		return
	default:
		return
	}
}

func CalcularEstoque(estoque int, retirada int) int {
	return estoque - retirada
}

func ListarProdutos(est *Estoque) {
	for i := 0; i < len(est.produtos); i++ {
		fmt.Println("Nome: ", est.produtos[i].nome)
		fmt.Println("Preço: ", est.produtos[i].preco)
		fmt.Println("Quantidade: ", est.produtos[i].quantidade)
		fmt.Println("-------------------------------")
	}
}

func ListarComprasCarrinho(carrinho *Carrinho) {
	for i := 0; i < len(carrinho.compras); i++ {
		fmt.Println("Nome: ", carrinho.compras[i].produto.nome)
		fmt.Println("Preço: ", carrinho.compras[i].valor)
		fmt.Println("Quantidade: ", carrinho.compras[i].quantidade)
		fmt.Println("-------------------------------")
	}
}

func AddProdutoCarinho(nome string, quantidade uint16, carrinho *Carrinho, estoque *Estoque) {
	var compra Compra
	for i := 0; i < len(estoque.produtos); i++ {
		if nome == estoque.produtos[i].nome {
			compra.produto = estoque.produtos[i]
			compra.quantidade = int(quantidade)
			compra.valor = (estoque.produtos[i].preco) * float32(quantidade)
			estoque.produtos[i].quantidade -= int(quantidade)
			carrinho.compras = append(carrinho.compras, compra) // func append add um valor ao final do Slice
			carrinho.valor += compra.valor
			carrinho.valor = float32(math.Floor(float64(carrinho.valor)*100) / 100) // Arredonda para baixo um número fracionário com 3 ou mais casa para duas casas
		}
	}
}

func ExcluiCompraCarrinho(nome string, carrinho *Carrinho) {
	var compra []Compra
	for i := 0; i < len(carrinho.compras); i++ {
		if nome == carrinho.compras[i].produto.nome {
			carrinho.valor -= carrinho.compras[i].valor
			carrinho.valor = float32(math.Floor(float64(carrinho.valor)*100) / 100)
		} else {
			compra = append(compra, carrinho.compras[i])
		}
	}
	carrinho.compras = compra
}

func ValidaNomeProduto(nome string, estoque *Estoque) bool {
	for i := 0; i < len(estoque.produtos); i++ {
		if nome == estoque.produtos[i].nome {
			return true
		}
	}
	return false
}

func ValidaQuantidadeProduto(nome string, quantidade uint16, estoque *Estoque) bool {
	for i := 0; i < len(estoque.produtos); i++ {
		if nome == estoque.produtos[i].nome {
			if estoque.produtos[i].quantidade >= int(quantidade) {
				return true
			}
			return false
		}
	}
	return false
}

func ValidaNomeCompra(nome string, carrinho *Carrinho) bool {
	for i := 0; i < len(carrinho.compras); i++ {
		if nome == carrinho.compras[i].produto.nome {
			return true
		}
	}
	return false
}

func ValidarPagamento(valorCompra float32, dinheiro float32) bool {
	valorCompra -= dinheiro
	valorCompra = float32(math.Floor(float64(valorCompra)*100) / 100)
	if valorCompra <= 0.0 {
		fmt.Println("Seu troco: ", (valorCompra * -1), "Reais")
		return true
	} else if valorCompra > 0.0 {
		for valorCompra > 0.0 {
			fmt.Println("Ainda falta: ", valorCompra, "Reais, favor completar o valor")
			fmt.Scan(&dinheiro)
			valorCompra -= dinheiro
			valorCompra = float32(math.Floor(float64(valorCompra)*100) / 100)
			if valorCompra < 0.0 {
				fmt.Println("Seu troco: ", (valorCompra * -1), "Reais")
			}
		}
		return true
	}
	return false
}
