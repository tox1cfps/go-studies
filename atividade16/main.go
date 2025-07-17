package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var chute int
	fmt.Println("\nGuess the Number!")
	fmt.Print("\nChute um número: ")
	fmt.Scan(&chute)
	var contador int
	contador = 1

	numero := rand.Intn(11)

jogo:
	for {
		if chute == numero {
			contador += 1
			fmt.Println("Parabéns, você acertou!")
			break jogo
		} else if chute < numero {
			contador += 1
			fmt.Println("Muito baixo!")
			fmt.Print("Chute um numero: ")
			fmt.Scan(&chute)
		} else if chute > numero {
			contador += 1
			fmt.Println("Muito alto!")
			fmt.Print("Chute um numero: ")
			fmt.Scan(&chute)
		}

		if contador == 7 {
			fmt.Printf("Você perdeu! o número era: %d", numero)
			return
		}
	}

	fmt.Printf("Você precisou de %d palpites para acertar!", contador)
}
