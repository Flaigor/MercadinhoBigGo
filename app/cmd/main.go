package main

import ( // forma recomendade de declarar vários imports
	"fmt" // biblioteca que tem as funcs do C
	"math"
	"os"
	"os/exec"

	"mercadinhoBigGo/domain/entities"
	"mercadinhoBigGo/domain/services"
)

var ( // também funciona com várias variáveis
	carrinho  entities.Carrinho
	estoque   entities.Estoque
	Cliente   entities.Cliente
	pCarrinho *entities.Carrinho = &carrinho
	pEstoque  *entities.Estoque  = &estoque
	pCliente  *entities.Cliente  = &Cliente
)

func main() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout

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

	// o número depois do tipo exemplo "int8", esse "8" representa o número de bits que essa váriavel vai usar
	var (
		op                int8 // int8 int16 int32 int64
		NomeProduto       string
		QuantidadeProduto uint16
		dinheiro          float32
	)

	var (
		pInt     *int8    = nil // nil é o NULL em GoLang
		pUint    *uint16  = nil
		pString  *string  = nil
		pFloat32 *float32 = nil
	)

	pInt = &op
	pString = &NomeProduto
	pUint = &QuantidadeProduto
	pFloat32 = &dinheiro

	fmt.Println("Nome do Cliente:")
	fmt.Scan(&Cliente.Nome)
	Inicializacao()

	for true { // while

		cmd.Run() // limpa tela no cmd

		carrinho.Valor = float32(math.Floor(float64(carrinho.Valor)*100) / 100)

		services.Saudacoes(Cliente.Nome)
		fmt.Println("Seu carrinho está em", carrinho.Valor, "Reais")
		services.MontaMenurPrincipal()

		fmt.Scan(pInt)

		switch op {
		case 1:
			services.ListarProdutos(pEstoque)
			fmt.Println("Qual produto quer adicionar ao carrinho: ")
			fmt.Scan(pString)
			if services.ValidaNomeProduto(*pString, pEstoque) {
				fmt.Println("Quantas unidades?")
				fmt.Scan(pUint)
				if services.ValidaQuantidadeProduto(*pString, *pUint, pEstoque) {
					services.AddProdutoCarinho(*pString, *pUint, pCarrinho, pEstoque)
				} else {
					fmt.Println("Quantidade inválida")
				}
			} else {
				fmt.Println("Produto inválido")
			}

			break

		case 2:
			services.ListarComprasCarrinho(pCarrinho)
			fmt.Println("Qual compra quer excluir do carrinho: ")
			fmt.Scan(pString)
			if services.ValidaNomeCompra(*pString, pCarrinho) {
				services.ExcluiCompraCarrinho(*pString, pCarrinho)
			} else {
				fmt.Println("Compra inválido")
			}

			break

		case 3:
			fmt.Println("Compras presentes no seu carrinho: ")
			services.ListarComprasCarrinho(pCarrinho)
			break

		case 4:
			fmt.Println("Compras presentes no seu carrinho: ")
			services.ListarComprasCarrinho(pCarrinho)
			fmt.Println("Deseja fechar o pedido? S/N: ")
			fmt.Scan(pString)
			if *pString == "S" || *pString == "s" {
				fmt.Println("Seu carrinho está em", carrinho.Valor, "Reais")
				fmt.Println("Quanto deseja pagar?: ")
				fmt.Scan(pFloat32)
				if services.ValidarPagamento(carrinho.Valor, *pFloat32) {
					fmt.Println("Agradecemos sua preferência, volte sempre!")
					return
				}
			}
			break

		case 0:
			return
		}
	}
	return
}

func Inicializacao() { // carregamento dos Produtos
	carrinho.Valor = 0.0
	carrinho.Cliente = Cliente
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
