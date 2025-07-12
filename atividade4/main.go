package main

import "fmt"

func main() {
	var tempoVivo int
	fmt.Print("Digite o ano em que nasceu: ")
	fmt.Scan(&tempoVivo)

	anoAtual := 2025

	diasTotais := (anoAtual - tempoVivo) * 365
	anos := diasTotais / 365
	meses := (diasTotais % 365) / 30
	dias := (diasTotais % 365) % 30

	fmt.Printf("Você vive há %d anos, %d meses e %d dias!", anos, meses, dias)
}
