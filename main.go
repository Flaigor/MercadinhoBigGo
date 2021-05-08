package main

import ( // forma recomendade de declarar vários imports
	"fmt" // biblioteca que tem as funcs do C
	"math"
	"os"
	"os/exec"
	//"rsc.io/quote" // $ go mod tidy
)

var ( // também funciona com várias variáveis
	carrinho  Carrinho
	estoque   Estoque
	cliente   Cliente
	pCarrinho *Carrinho = &carrinho
	pEstoque  *Estoque  = &estoque
	pCliente  *Cliente  = &cliente
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
		nomeProduto       string
		quantidadeProduto uint16
		dinheiro          float32
	)

	var (
		pInt     *int8    = nil // nil é o NULL em GoLang
		pUint    *uint16  = nil
		pString  *string  = nil
		pFloat32 *float32 = nil
	)

	pInt = &op
	pString = &nomeProduto
	pUint = &quantidadeProduto
	pFloat32 = &dinheiro

	fmt.Println("Nome do Cliente:")
	fmt.Scan(&cliente.nome)
	Inicializacao()

	for true == true { // while

		cmd.Run() // limpa tela no cmd

		carrinho.valor = float32(math.Floor(float64(carrinho.valor)*100) / 100)

		Saudacoes(cliente.nome)
		fmt.Println("Seu carrinho está em", carrinho.valor, "Reais")
		MontaMenurPrincipal()

		fmt.Scan(pInt)

		switch op {
		case 1:
			ListarProdutos(pEstoque)
			fmt.Println("Qual produto quer adicionar ao carrinho: ")
			fmt.Scan(pString)
			if ValidaNomeProduto(*pString, pEstoque) {
				fmt.Println("Quantas unidades?")
				fmt.Scan(pUint)
				if ValidaQuantidadeProduto(*pString, *pUint, pEstoque) {
					AddProdutoCarinho(*pString, *pUint, pCarrinho, pEstoque)
				} else {
					fmt.Println("Quantidade inválida")
				}
			} else {
				fmt.Println("Produto inválido")
			}

			break

		case 2:
			ListarComprasCarrinho(pCarrinho)
			fmt.Println("Qual compra quer excluir do carrinho: ")
			fmt.Scan(pString)
			if ValidaNomeCompra(*pString, pCarrinho) {
				ExcluiCompraCarrinho(*pString, pCarrinho)
			} else {
				fmt.Println("Compra inválido")
			}

			break

		case 3:
			fmt.Println("Compras presentes no seu carrinho: ")
			ListarComprasCarrinho(pCarrinho)
			break

		case 4:
			fmt.Println("Compras presentes no seu carrinho: ")
			ListarComprasCarrinho(pCarrinho)
			fmt.Println("Deseja fechar o pedido? S/N: ")
			fmt.Scan(pString)
			if *pString == "S" || *pString == "s" {
				fmt.Println("Seu carrinho está em", carrinho.valor, "Reais")
				fmt.Println("Quanto deseja pagar?: ")
				fmt.Scan(pFloat32)
				if ValidarPagamento(carrinho.valor, *pFloat32) {
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

func Inicializacao() { // carregamento dos produtos
	carrinho.valor = 0.0
	carrinho.cliente = cliente
	carrinho.compras = nil

	var (
		carne        Produto
		peixe        Produto
		arroz        Produto
		feijao       Produto
		suco         Produto
		batata       Produto
		queijo       Produto
		refrigerante Produto
		frango       Produto
		leite        Produto
	)

	carne.nome = "Carne"
	peixe.nome = "Peixe"
	arroz.nome = "Arroz"
	feijao.nome = "Feijão"
	suco.nome = "Suco"
	batata.nome = "Batata"
	queijo.nome = "Queijo"
	refrigerante.nome = "Refrigerante"
	frango.nome = "Frango"
	leite.nome = "Leite"

	carne.preco = 57.99
	peixe.preco = 43.99
	arroz.preco = 15.99
	feijao.preco = 7.99
	suco.preco = 5.98
	batata.preco = 4.30
	queijo.preco = 1.50
	refrigerante.preco = 7.00
	frango.preco = 12.99
	leite.preco = 6.00

	carne.quantidade = 100
	peixe.quantidade = 25
	arroz.quantidade = 30
	feijao.quantidade = 50
	suco.quantidade = 300
	batata.quantidade = 1300
	queijo.quantidade = 70
	refrigerante.quantidade = 150
	frango.quantidade = 100
	leite.quantidade = 250

	estoque.produtos = append(estoque.produtos, carne)
	estoque.produtos = append(estoque.produtos, peixe)
	estoque.produtos = append(estoque.produtos, arroz)
	estoque.produtos = append(estoque.produtos, feijao)
	estoque.produtos = append(estoque.produtos, suco)
	estoque.produtos = append(estoque.produtos, batata)
	estoque.produtos = append(estoque.produtos, queijo)
	estoque.produtos = append(estoque.produtos, refrigerante)
	estoque.produtos = append(estoque.produtos, frango)
	estoque.produtos = append(estoque.produtos, leite)
}
