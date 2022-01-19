// +build unit
package services_test

import (
	"encoding/json"
	"fmt"
	"math"
	"mercadinhoBigGo/domain/entities"
	"mercadinhoBigGo/domain/services"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TEST TYPE: SIMPLE
func TestCalculaQuadradoECubo(t *testing.T) {
	var tests = []struct {
		value                rune
		wantSquare, wantCube int
	}{
		{0, 0, 0},
		{1, 1, 1},
		{2, 4, 8},
		{3, 9, 27},
		{4, 16, 64},
		{5, 25, 125},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.value)
		t.Run(testname, func(t *testing.T) {
			square, cube := services.CalculaQuadradoECubo(tt.value)
			if square != tt.wantSquare {
				t.Errorf("got %d, want %d", square, tt.wantSquare)
			}
			if cube != tt.wantCube {
				t.Errorf("got %d, want %d", square, tt.wantCube)
			}
		})
	}
}

func TestCalcularDoisValores(t *testing.T) {
	var tests = []struct {
		firstValue, secondValue float32
		operation               uint8
		want                    float32
	}{
		//sum
		{0, 0, 1, 0},
		{0, 1, 1, 1},
		{15, 35, 1, 50},
		{-15, 35, 1, 20},
		{15, -35, 1, -20},
		{-15, -35, 1, -50},
		//subtraction
		{0, 0, 2, 0},
		{0, 1, 2, -1},
		{15, 35, 2, -20},
		{-15, 35, 2, -50},
		{15, -35, 2, 50},
		{-15, -35, 2, 20},
		//multiplication
		{0, 0, 3, 0},
		{0, 1, 3, 0},
		{2, 35, 3, 70},
		{-2, 35, 3, -70},
		{2, -35, 3, -70},
		{-2, -35, 3, 70},
		{2.6, 2, 3, 5.2},
		//division
		{0, 1, 4, 0},
		{8, 2, 4, 4},
		{-8, 2, 4, -4},
		{8, -2, 4, -4},
		{-8, -2, 4, 4},
		{2.6, 2, 4, 1.3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f,%f,%d", tt.firstValue, tt.secondValue, tt.operation)
		t.Run(testname, func(t *testing.T) {
			response := services.CalcularDoisValores(tt.firstValue, tt.secondValue, tt.operation)
			if response != tt.want {
				t.Errorf("got %f, want %f", response, tt.want)
			}
		})
	}
}

func TestCalcularEstoque(t *testing.T) {
	var tests = []struct {
		firstValue, secondValue int
		want                    int
	}{
		{0, 0, 0},
		{1, 1, 0},
		{2, 4, -2},
		{9, 8, 1},
		{16, 4, 12},
		{125, 25, 100},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.firstValue, tt.secondValue)
		t.Run(testname, func(t *testing.T) {
			inventory := services.CalcularEstoque(tt.firstValue, tt.secondValue)
			if inventory != tt.want {
				t.Errorf("got %d, want %d", inventory, tt.want)
			}
		})
	}
}

func TestAddProdutoCarinho(t *testing.T) {
	var tests = []struct {
		testName     string
		addProducts  []string
		oldProducts  []entities.Produto
		wantQuantity int
		wantPrice    float32
	}{
		{
			"OneItem",
			[]string{
				"Batata",
			},
			[]entities.Produto{},
			1, 4.3,
		},
		{
			"TwoEqualItems",
			[]string{
				"Batata",
				"Batata",
			},
			[]entities.Produto{},
			2, 8.6,
		},
		{
			"TwoDifferentItems",
			[]string{
				"Batata",
				"Leite",
			},
			[]entities.Produto{},
			2, 10.3,
		},
		{
			"OneRepeatedItem",
			[]string{
				"Batata",
			},
			[]entities.Produto{
				AuxGeraProdutos("Batata"),
			},
			2, 8.6,
		},
		{
			"OneItemInCartWithTwoItems",
			[]string{
				"Batata",
			},
			[]entities.Produto{
				AuxGeraProdutos("Batata"),
				AuxGeraProdutos("Leite"),
			},
			3, 14.6,
		},
	}

	for _, tt := range tests {
		compras := []entities.Compra{}
		cartValue := float32(math.Floor(float64(0)*100) / 100)
		for _, item := range tt.oldProducts {
			purchase := entities.Compra{item, 1, item.Preco}
			compras = append(compras, purchase)
			cartValue += float32(math.Floor(float64(item.Preco)*100) / 100)
		}
		carrinho := entities.Carrinho{entities.Cliente{"Marcos"}, compras, cartValue}
		estoque := entities.Estoque{[]entities.Produto{AuxGeraProdutos("Batata"), AuxGeraProdutos("Leite")}}

		testname := fmt.Sprintf("%s", tt.testName)
		t.Run(testname, func(t *testing.T) {
			for _, item := range tt.addProducts {
				services.AddProdutoCarinho(item, 1, &carrinho, &estoque)
			}
			if len(carrinho.Compras) != tt.wantQuantity {
				t.Errorf("quantity got %d, quantity want %d", len(carrinho.Compras), tt.wantQuantity)
			}
			if carrinho.Valor != tt.wantPrice {
				t.Errorf("value got %f, value want %f", carrinho.Valor, tt.wantPrice)
			}
		})
	}
}

func TestExcluiCompraCarrinho(t *testing.T) {
	var tests = []struct {
		testName       string
		removeProducts []string
		oldProducts    []entities.Produto
		wantProducts   []entities.Produto
		wantQuantity   int
		wantPrice      float32
	}{
		{
			"EmptyCart",
			[]string{
				"Batata",
			},
			[]entities.Produto{},
			[]entities.Produto{},
			0,
			0.,
		},
		{
			"Simple",
			[]string{
				"Batata",
			},
			[]entities.Produto{
				AuxGeraProdutos("Batata"),
			},
			[]entities.Produto{},
			0,
			0.,
		},
		{
			"TwoEqualItensOnCart",
			[]string{
				"Batata",
			},
			[]entities.Produto{
				AuxGeraProdutos("Batata"),
				AuxGeraProdutos("Batata"),
			},
			[]entities.Produto{},
			0,
			0.,
		},
		{
			"TwoDifferentItemsOnCart",
			[]string{
				"Batata",
			},
			[]entities.Produto{
				AuxGeraProdutos("Batata"),
				AuxGeraProdutos("Leite"),
			},
			[]entities.Produto{
				AuxGeraProdutos("Leite"),
			},
			1,
			6.,
		},
		{
			"RemovingTwoWithOneLeft",
			[]string{
				"Batata",
				"Leite",
			},
			[]entities.Produto{
				AuxGeraProdutos("Batata"),
				AuxGeraProdutos("Leite"),
				AuxGeraProdutos("Carne"),
			},
			[]entities.Produto{
				AuxGeraProdutos("Carne"),
			},
			1,
			57.99,
		},
		{
			"RemovingTwoWithDuplicationAndThreeLeft",
			[]string{
				"Batata",
				"Leite",
			},
			[]entities.Produto{
				AuxGeraProdutos("Batata"),
				AuxGeraProdutos("Suco"),
				AuxGeraProdutos("Batata"),
				AuxGeraProdutos("Carne"),
				AuxGeraProdutos("Leite"),
				AuxGeraProdutos("Carne"),
			},
			[]entities.Produto{
				AuxGeraProdutos("Suco"),
				AuxGeraProdutos("Carne"),
				AuxGeraProdutos("Carne"),
			},
			3,
			121.96,
		},
	}

	for _, tt := range tests {
		compras := []entities.Compra{}
		cartValue := float32(math.Floor(float64(0)*100) / 100)
		for _, item := range tt.oldProducts {
			purchase := entities.Compra{item, 1, item.Preco}
			compras = append(compras, purchase)
			cartValue += float32(math.Floor(float64(item.Preco)*100) / 100)
		}
		carrinho := entities.Carrinho{entities.Cliente{"Marcos"}, compras, cartValue}

		testname := fmt.Sprintf("%s", tt.testName)
		t.Run(testname, func(t *testing.T) {
			for _, item := range tt.removeProducts {
				services.ExcluiCompraCarrinho(item, &carrinho)
			}
			if len(carrinho.Compras) != tt.wantQuantity {
				t.Errorf("quantity got %d, quantity want %d", len(carrinho.Compras), tt.wantQuantity)
			}
			if carrinho.Valor != tt.wantPrice {
				t.Errorf("value got %f, value want %f", carrinho.Valor, tt.wantPrice)
			}
		})
	}
}

// TEST TYPE: MOCKED HTTP REQUEST

func TestGetHostFromPost(t *testing.T) {
	var tests = []struct {
		testName string
		want     string
	}{
		{"Teste 1", "Carne"},
		{"Teste 2", "Suco"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			response := fmt.Sprintf(`{"Nome":"%s","Quantidade":5,"Preco":4.3}`, tt.want)
			httpmock.RegisterResponder(http.MethodPost, "https://httpbin.org/post",
				httpmock.NewStringResponder(200, response))

			resp := services.GetHostFromPost()

			product := entities.Produto{}

			json.Unmarshal(resp, &product)

			if product.Nome != tt.want {
				t.Errorf("got %s, want %s", product.Nome, tt.want)
			}
		})
	}
}

// TEST TYPE: EXAMPLE

func ExampleListarProdutos() {
	estoque := entities.Estoque{[]entities.Produto{AuxGeraProdutos("Batata"), AuxGeraProdutos("Leite")}}
	services.ListarProdutos(&estoque)
	// Output: Nome:  Batata
	// Preço:  4.3
	// Quantidade:  1300
	// -------------------------------
	// Nome:  Leite
	// Preço:  6
	// Quantidade:  250
	// -------------------------------
}

func ExampleListarComprasCarrinho() {
	compra := entities.Compra{AuxGeraProdutos("Batata"), 2, 8.6}
	carrinho := entities.Carrinho{entities.Cliente{"Marcos"}, []entities.Compra{compra}, 6.0}
	services.ListarComprasCarrinho(&carrinho)
	// Output: Nome:  Batata
	// Preço:  8.6
	// Quantidade:  2
	// -------------------------------
}

// TEST TYPE: AUXILIARY

func AuxGeraProdutos(nome string) entities.Produto {
	// if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Carne") {
	if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Carne")) {
		return entities.Produto{
			Nome:       "Carne",
			Preco:      57.99,
			Quantidade: 100,
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Peixe")) {
		return entities.Produto{
			Nome:       "Peixe",
			Preco:      43.99,
			Quantidade: 25,
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Arroz")) {
		return entities.Produto{
			Preco:      15.99,
			Quantidade: 30,
			Nome:       "Arroz",
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Feijão")) {
		return entities.Produto{
			Nome:       "Feijão",
			Preco:      7.99,
			Quantidade: 50,
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Suco")) {
		return entities.Produto{
			Nome:       "Suco",
			Preco:      5.98,
			Quantidade: 300,
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Batata")) {
		return entities.Produto{
			Nome:       "Batata",
			Preco:      4.30,
			Quantidade: 1300,
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Queijo")) {
		return entities.Produto{
			Nome:       "Queijo",
			Preco:      1.50,
			Quantidade: 70,
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Refrigerante")) {
		return entities.Produto{
			Nome:       "Refrigerante",
			Preco:      7.00,
			Quantidade: 150,
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Frango")) {
		return entities.Produto{
			Nome:       "Frango",
			Preco:      12.99,
			Quantidade: 100,
		}
	} else if strings.EqualFold(strings.ToLower(nome), strings.ToLower("Leite")) {
		return entities.Produto{
			Nome:       "Leite",
			Preco:      6.00,
			Quantidade: 250,
		}
	} else {
		return entities.Produto{}
	}
}
