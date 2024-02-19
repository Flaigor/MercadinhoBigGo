package main

import ( // forma recomendade de declarar vários imports
	"fmt" // biblioteca que tem as funcs do C
	"math"
	"os"
	"os/exec"

	"mercadinhoBigGo/domain/controllers"
	"mercadinhoBigGo/domain/entities"
)

var ( // também funciona com várias variáveis
	carrinho entities.Carrinho
	estoque  entities.Estoque
	cliente  entities.Cliente
)

func main() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout

	// o número depois do tipo exemplo "int8", esse "8" representa o número de bits que essa váriavel vai usar
	var (
		op   int8        // int8 int16 int32 int64
		pInt *int8 = nil // nil é o NULL em GoLang
	)

	pInt = &op

	fmt.Println("Nome do Cliente:")
	fmt.Scan(&cliente.Nome)
	controller := controllers.Controller{}
	controller.Inicializacao(&estoque, &carrinho)

	for true { // while

		cmd.Run() // limpa tela no cmd

		carrinho.Valor = float32(math.Floor(float64(carrinho.Valor)*100) / 100)

		controller.Saudacoes(cliente.Nome)
		fmt.Println("Seu carrinho está em", carrinho.Valor, "Reais")
		controller.MontaMenurPrincipal()

		fmt.Scan(pInt)

		finished, _ := controller.Process(pInt, &estoque, &carrinho)
		if finished {
			break
		}
	}
	return
}
