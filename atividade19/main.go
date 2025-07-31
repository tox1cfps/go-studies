package main

import "fmt"

func eCrianca(i int) bool {
	return i >= 0 && i < 12
}

func eAdolescente(i int) bool {
	return i >= 12 && i <= 17
}

func eAdulto(i int) bool {
	return i >= 18 && i <= 59
}

func eIdoso(i int) bool {
	return i >= 60
}

func checarIdade(i int) (bool, string) {

	if i < 0 {
		return false, ""
	}

	if eCrianca(i) {
		return true, "Criança"
	} else if eAdolescente(i) {
		return true, "Adolescente"
	} else if eAdulto(i) {
		return true, "Adulto"
	} else if eIdoso(i) {
		return true, "Idoso"
	}
	return false, ""
}

func main() {
	var criancas, adolescentes, adultos, idosos int

	for {
		var idade int
		fmt.Print("Insira a idade: ")
		fmt.Scan(&idade)

		valida, grupo := checarIdade(idade)
		if !valida {
			break
		}

		fmt.Println("Classificação: ", grupo)
		switch grupo {
		case "Criança":
			criancas++
		case "Adolescente":
			adolescentes++
		case "Adulto":
			adultos++
		case "Idoso":
			idosos++
		}
	}
	fmt.Println("\nResumo Final")
	fmt.Println("Crianças: ", criancas)
	fmt.Println("Adolescentes: ", adolescentes)
	fmt.Println("Adultos: ", adultos)
	fmt.Println("Idosos: ", idosos)
}
