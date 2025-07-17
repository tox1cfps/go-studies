package main

import "fmt"

func main() {
	var candidato1 int
	var candidato2 int
	var candidato3 int
	var nulo int

votacao:
	for {
		var escolha int
		fmt.Println("\nMenu de votaçao!")
		fmt.Print("\n1. Candidato A")
		fmt.Print("\n2. Candidato B")
		fmt.Print("\n3. Candidato C")
		fmt.Print("\n0. Sair")
		fmt.Print("\nEscolha seu candidato: ")
		fmt.Scan(&escolha)

		switch escolha {
		case 0:
			fmt.Println("Obrigado por votar!")
			break votacao
		case 1:
			candidato1 += 1
			fmt.Println("\nObrigado por votar!")
		case 2:
			candidato2 += 1
			fmt.Println("\nObrigado por votar!")
		case 3:
			candidato3 += 1
			fmt.Println("\nObrigado por votar!")
		default:
			nulo += 1
			fmt.Println("\nObrigado por votar!")
		}
	}

	fmt.Println("\nApuração dos Votos")
	fmt.Printf("Candidato A recebeu %d votos\n", candidato1)
	fmt.Printf("Candidato B recebeu %d votos\n", candidato2)
	fmt.Printf("Candidato C recebeu %d votos\n", candidato3)
	fmt.Printf("Votos nulos: %d\n", nulo)

	if candidato1 > candidato2 && candidato1 > candidato3 {
		fmt.Println("\nCandidato A venceu a eleição!")
	} else if candidato2 > candidato1 && candidato2 > candidato3 {
		fmt.Println("\nCandidato B venceu a eleição!")
	} else if candidato3 > candidato1 && candidato3 > candidato2 {
		fmt.Println("\nCandidato C venceu a eleição!")
	} else if candidato1 == candidato2 || candidato2 == candidato3 || candidato3 == candidato1 {
		fmt.Println("Candidatos empatados iram para uma segunda eleição")
	} else if nulo > candidato1 && nulo > candidato2 && nulo > candidato3 {
		fmt.Println("Votação cancelada")
	}
}
