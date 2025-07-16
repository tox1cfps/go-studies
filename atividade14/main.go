package main

import "fmt"

func main() {
	var Tempo float64
	fmt.Print("Insira em quanto tempo deseja realizar a viagem: ")
	fmt.Scan(&Tempo)
	var VelocidadeMedia float64
	fmt.Print("Digite a velocidade média em que fará essa viagem: ")
	fmt.Scan(&VelocidadeMedia)

	distancia := Tempo * VelocidadeMedia
	litrosUsados := distancia / 12

	fmt.Printf("%0.2f Litros serão gastos em %.02f Kilometros rodados em uma média de %.02f KM/H em %.02f Horas", litrosUsados, distancia, VelocidadeMedia, Tempo)
}
