package controllers

import (
	"context"
	"fmt"
	"mercadinhoBigGo/domain/entities"
	"mercadinhoBigGo/domain/services"
)

type Controller struct {
	Ctx context.Context
}

func (controller Controller) Saudacoes(Nome string) { // declaração base de uma func
	fmt.Println("Olá", Nome, ", seja bem vindo!")
}

func (controller Controller) MontaMenurPrincipal() {
	fmt.Println("|| 1 - Adicionar Produto ao carrinho\n|| 2 - Excluir prduto do carrinho")
	fmt.Println("|| 3 - Consultar Carrinho\n|| 4 - Fechar Pedido\n|| 0 - Fazer Logoff")
}

func (controller Controller) Inicializacao(estoque *entities.Estoque, carrinho *entities.Carrinho) { // carregamento dos Produtos
	carrinho.Valor = 0.0
	carrinho.Compras = nil

	var (
		carne        entities.Produto
		peixe        entities.Produto
		arroz        entities.Produto
		feijao       entities.Produto
		suco         entities.Produto
		batata       entities.Produto
		queijo       entities.Produto
		refrigerante entities.Produto
		frango       entities.Produto
		leite        entities.Produto
	)

	carne.Nome = "Carne"
	peixe.Nome = "Peixe"
	arroz.Nome = "Arroz"
	feijao.Nome = "Feijão"
	suco.Nome = "Suco"
	batata.Nome = "Batata"
	queijo.Nome = "Queijo"
	refrigerante.Nome = "Refrigerante"
	frango.Nome = "Frango"
	leite.Nome = "Leite"

	carne.Preco = 57.99
	peixe.Preco = 43.99
	arroz.Preco = 15.99
	feijao.Preco = 7.99
	suco.Preco = 5.98
	batata.Preco = 4.30
	queijo.Preco = 1.50
	refrigerante.Preco = 7.00
	frango.Preco = 12.99
	leite.Preco = 6.00

	carne.Quantidade = 100
	peixe.Quantidade = 25
	arroz.Quantidade = 30
	feijao.Quantidade = 50
	suco.Quantidade = 300
	batata.Quantidade = 1300
	queijo.Quantidade = 70
	refrigerante.Quantidade = 150
	frango.Quantidade = 100
	leite.Quantidade = 250

	estoque.Produtos = append(estoque.Produtos, carne)
	estoque.Produtos = append(estoque.Produtos, peixe)
	estoque.Produtos = append(estoque.Produtos, arroz)
	estoque.Produtos = append(estoque.Produtos, feijao)
	estoque.Produtos = append(estoque.Produtos, suco)
	estoque.Produtos = append(estoque.Produtos, batata)
	estoque.Produtos = append(estoque.Produtos, queijo)
	estoque.Produtos = append(estoque.Produtos, refrigerante)
	estoque.Produtos = append(estoque.Produtos, frango)
	estoque.Produtos = append(estoque.Produtos, leite)
}

func (controller Controller) Process(op *int8, estoque *entities.Estoque, carrinho *entities.Carrinho) (bool, error) {

	var (
		quantidade  uint16  = 0
		NomeProduto string  = ""
		dinheiro    float32 = 0
	)

	switch *op {
	case 1:
		services.ListarProdutos(estoque)
		fmt.Println("Qual produto quer adicionar ao carrinho: ")
		fmt.Scan(&NomeProduto)
		if services.ValidaNomeProduto(NomeProduto, estoque) {
			fmt.Println("Quantas unidades?")
			fmt.Scan(&quantidade)
			if services.ValidaQuantidadeProduto(NomeProduto, quantidade, estoque) {
				services.AddProdutoCarinho(NomeProduto, quantidade, carrinho, estoque)
			} else {
				fmt.Println("Quantidade inválida")
			}
		} else {
			fmt.Println("Produto inválido")
		}

		break

	case 2:
		services.ListarComprasCarrinho(carrinho)
		fmt.Println("Qual compra quer excluir do carrinho: ")
		fmt.Scan(&NomeProduto)
		if services.ValidaNomeCompra(NomeProduto, carrinho) {
			services.ExcluiCompraCarrinho(NomeProduto, carrinho)
		} else {
			fmt.Println("Compra inválido")
		}

		break

	case 3:
		fmt.Println("Compras presentes no seu carrinho: ")
		services.ListarComprasCarrinho(carrinho)
		break

	case 4:
		fmt.Println("Compras presentes no seu carrinho: ")
		services.ListarComprasCarrinho(carrinho)
		fmt.Println("Deseja fechar o pedido? S/N: ")
		fmt.Scan(&NomeProduto)
		if NomeProduto == "S" || NomeProduto == "s" {
			fmt.Println("Seu carrinho está em", carrinho.Valor, "Reais")
			fmt.Println("Quanto deseja pagar?: ")
			fmt.Scan(&dinheiro)
			if services.ValidarPagamento(carrinho.Valor, dinheiro) {
				fmt.Println("Agradecemos sua preferência, volte sempre!")
				return true, nil
			}
		}
		break

	case 0:
		return true, nil
	}

	return false, nil
}

// Demonstração de como receber um função de retorno duplo
/*
	var quadrado, cubo int

	quadrado, cubo = CalculaQuadradoECubo(2)
	fmt.Println(quadrado, cubo)

	return

*/

// Demonstração de criação de um Slice usando make, como altera sua len e cap e alteração entre heranças
/*
	var aux []int
	array := make([]int, 0, 5) // func maker cria um array, nesse caso é um Slice com cap = 5

	fmt.Println("Array len:", len(array), "Array cap:", cap(array))
	fmt.Println("Aux len:", len(aux), "Z cap:", cap(aux))

	for i := 0; i < cap(array); i++ {
		array = append(array, i)
	}

	aux = array

	fmt.Println("Array:", array)
	fmt.Println("Array len:", len(array), "Array cap:", cap(array))
	fmt.Println("Aux len:", len(aux), "Aux cap:", cap(aux))

	array = aux[2:5] // Array agora é um pedaço do Aux

	fmt.Println("Array len:", len(array), "Array cap:", cap(array))
	fmt.Println("Aux len:", len(aux), "Aux cap:", cap(aux))

	aux[2] = 1 // Aux altera Array

	fmt.Println("Array:", array)
	fmt.Println("Aux:", aux)

	array[0] = 2 // Array altera Aux

	fmt.Println("Array:", array)
	fmt.Println("Aux:", aux)

	return

*/
