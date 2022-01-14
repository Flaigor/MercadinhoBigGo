package services_test

import (
	"encoding/json"
	"fmt"
	"math"
	"mercadinhoBigGo/domain/entities"
	"mercadinhoBigGo/domain/services"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TEST TYPE: SIMPLE TEST

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
		wantValue    float32
	}{
		{
			"OneItem",
			[]string{
				"Batata",
			},
			[]entities.Produto{},
			1, 3.0,
		},
		{
			"TwoEqualItems",
			[]string{
				"Batata",
				"Batata",
			},
			[]entities.Produto{},
			2, 6.0,
		},
		{
			"TwoDifferentItems",
			[]string{
				"Batata",
				"Guaraná",
			},
			[]entities.Produto{},
			2, 5.5,
		},
		{
			"OneRepeatedItem",
			[]string{
				"Batata",
			},
			[]entities.Produto{
				entities.Produto{"Batata", 7, 3.0},
			},
			2, 6.0,
		},
		{
			"OneItemInCartWithTwoItems",
			[]string{
				"Batata",
			},
			[]entities.Produto{
				entities.Produto{"Batata", 7, 3.0},
				entities.Produto{"Guaraná", 12, 2.5},
			},
			3, 8.5,
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
		estoque := entities.Estoque{[]entities.Produto{entities.Produto{"Batata", 7, 3.0}, entities.Produto{"Guaraná", 12, 2.5}}}

		testname := fmt.Sprintf("%s", tt.testName)
		t.Run(testname, func(t *testing.T) {
			for _, item := range tt.addProducts {
				services.AddProdutoCarinho(item, 1, &carrinho, &estoque)
			}
			if len(carrinho.Compras) != tt.wantQuantity {
				t.Errorf("quantity got %d, quantity want %d", len(carrinho.Compras), tt.wantQuantity)
			}
			if carrinho.Valor != tt.wantValue {
				t.Errorf("value got %f, value want %f", carrinho.Valor, tt.wantValue)
			}
		})
	}
}

func TestGetHostFromPost(t *testing.T) {
	var tests = []struct {
		testName string
		want     string
	}{
		{"Teste 1", "Tomate"},
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
	estoque := entities.Estoque{[]entities.Produto{entities.Produto{"Batata", 7, 3.0}, entities.Produto{"Guaraná", 12, 2.5}}}
	services.ListarProdutos(&estoque)
	// Output: Nome:  Batata
	// Preço:  3
	// Quantidade:  7
	// -------------------------------
	// Nome:  Guaraná
	// Preço:  2.5
	// Quantidade:  12
	// -------------------------------
}

func ExampleListarComprasCarrinho() {
	compra := entities.Compra{entities.Produto{"Batata", 7, 3.0}, 2, 6.0}
	carrinho := entities.Carrinho{entities.Cliente{"Marcos"}, []entities.Compra{compra}, 6.0}
	services.ListarComprasCarrinho(&carrinho)
	// Output: Nome:  Batata
	// Preço:  6
	// Quantidade:  2
	// -------------------------------
}
