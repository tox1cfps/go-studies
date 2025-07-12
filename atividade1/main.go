package main

import "fmt"

func main() {
	var preçoproduto float64
	fmt.Print("Digite o preço do produto: ")
	fmt.Scan(&preçoproduto)

	// loop com tipos de pagamentos
	for {
		var metodo int
		fmt.Println("\nMenu de Pagamento")
		fmt.Print("\n1. Dinheiro a vista ou PIX - %15 Off\n")
		fmt.Print("\n2. A vista no cartão de crédito - %10 Off\n")
		fmt.Print("\n3. Parcelado no cartão em até duas vezes (Sem Juros)\n")
		fmt.Print("\n4. Parcelado no cartão em até três vezes ou mais - +10% Juros\n")
		fmt.Print("Escolha o método de pagamento: ")
		fmt.Scan(&metodo)

		var ajuste float64

		switch metodo {
		case 1:
			ajuste = preçoproduto * 0.85
			fmt.Printf("Preço a ser pago: %.02f", ajuste)
			return
		case 2:
			ajuste = preçoproduto * 0.90
			fmt.Printf("Preço a ser pago: %.02f", ajuste)
			return
		case 3:
			ajuste = preçoproduto * 1.00
			fmt.Printf("Preço a ser pago: %.02f", ajuste)
			return
		case 4:
			ajuste = preçoproduto * 1.10
			fmt.Printf("Preço a ser pago: %.02f", ajuste)
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
