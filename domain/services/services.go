package services

import (
	"fmt"
	"math"

	"mercadinhoBigGo/domain/entities"
)

func Saudacoes(Nome string) { // declaração base de uma func
	fmt.Println("Olá", Nome, ", seja bem vindo!")
}

func MontaMenurPrincipal() {
	fmt.Println("|| 1 - Adicionar Produto ao carrinho\n|| 2 - Excluir prduto do carrinho")
	fmt.Println("|| 3 - Consultar Carrinho\n|| 4 - Fechar Pedido\n|| 0 - Fazer Logoff")
}

func CalculaQuadradoECubo(x rune) (int, int) { // É possivel retornar mais de um Valor em uma func
	// rune = int32 porém precisa ser convertido, caso não seja int32
	return int((x * x)), int((x * x * x))
}

func CalcularDoisValores(x float32, y float32, op uint8) (resp float32) { // Valores de retorno Nomeados não são recomendados para funcs muito grandes
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

func ListarProdutos(est *entities.Estoque) {
	for i := 0; i < len(est.Produtos); i++ {
		fmt.Println("Nome: ", est.Produtos[i].Nome)
		fmt.Println("Preço: ", est.Produtos[i].Preco)
		fmt.Println("Quantidade: ", est.Produtos[i].Quantidade)
		fmt.Println("-------------------------------")
	}
}

func ListarComprasCarrinho(carrinho *entities.Carrinho) {
	for i := 0; i < len(carrinho.Compras); i++ {
		fmt.Println("Nome: ", carrinho.Compras[i].Produto.Nome)
		fmt.Println("Preço: ", carrinho.Compras[i].Valor)
		fmt.Println("Quantidade: ", carrinho.Compras[i].Quantidade)
		fmt.Println("-------------------------------")
	}
}

func AddProdutoCarinho(Nome string, Quantidade uint16, carrinho *entities.Carrinho, estoque *entities.Estoque) {
	var compra entities.Compra
	for i := 0; i < len(estoque.Produtos); i++ {
		if Nome == estoque.Produtos[i].Nome {
			compra.Produto = estoque.Produtos[i]
			compra.Quantidade = int(Quantidade)
			compra.Valor = (estoque.Produtos[i].Preco) * float32(Quantidade)
			estoque.Produtos[i].Quantidade -= int(Quantidade)
			carrinho.Compras = append(carrinho.Compras, compra) // func append add um Valor ao final do Slice
			carrinho.Valor += compra.Valor
			carrinho.Valor = float32(math.Floor(float64(carrinho.Valor)*100) / 100) // Arredonda para baixo um número fracionário com 3 ou mais casa para duas casas
		}
	}
}

func ExcluiCompraCarrinho(Nome string, carrinho *entities.Carrinho) {
	var compra []entities.Compra
	for i := 0; i < len(carrinho.Compras); i++ {
		if Nome == carrinho.Compras[i].Produto.Nome {
			carrinho.Valor -= carrinho.Compras[i].Valor
			carrinho.Valor = float32(math.Floor(float64(carrinho.Valor)*100) / 100)
		} else {
			compra = append(compra, carrinho.Compras[i])
		}
	}
	carrinho.Compras = compra
}

func ValidaNomeProduto(Nome string, estoque *entities.Estoque) bool {
	for i := 0; i < len(estoque.Produtos); i++ {
		if Nome == estoque.Produtos[i].Nome {
			return true
		}
	}
	return false
}

func ValidaQuantidadeProduto(Nome string, Quantidade uint16, estoque *entities.Estoque) bool {
	for i := 0; i < len(estoque.Produtos); i++ {
		if Nome == estoque.Produtos[i].Nome {
			if estoque.Produtos[i].Quantidade >= int(Quantidade) {
				return true
			}
			return false
		}
	}
	return false
}

func ValidaNomeCompra(Nome string, carrinho *entities.Carrinho) bool {
	for i := 0; i < len(carrinho.Compras); i++ {
		if Nome == carrinho.Compras[i].Produto.Nome {
			return true
		}
	}
	return false
}

func ValidarPagamento(ValorCompra float32, dinheiro float32) bool {
	ValorCompra -= dinheiro
	ValorCompra = float32(math.Floor(float64(ValorCompra)*100) / 100)
	if ValorCompra <= 0.0 {
		fmt.Println("Seu troco: ", (ValorCompra * -1), "Reais")
		return true
	} else if ValorCompra > 0.0 {
		for ValorCompra > 0.0 {
			fmt.Println("Ainda falta: ", ValorCompra, "Reais, favor completar o Valor")
			fmt.Scan(&dinheiro)
			ValorCompra -= dinheiro
			ValorCompra = float32(math.Floor(float64(ValorCompra)*100) / 100)
			if ValorCompra < 0.0 {
				fmt.Println("Seu troco: ", (ValorCompra * -1), "Reais")
			}
		}
		return true
	}
	return false
}