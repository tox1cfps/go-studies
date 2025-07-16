package main

import "fmt"

func main() {
	var HoraAula float64
	fmt.Print("Insira o valor da Hora aula que recebe: ")
	fmt.Scan(&HoraAula)
	var NumeroAulas float64
	fmt.Print("Insira o número de aulas que foram dadas no mês: ")
	fmt.Scan(&NumeroAulas)

	resultado := HoraAula * NumeroAulas
	INSS := resultado * 0.075

	if resultado <= 1518.00 {
		fmt.Printf("Salário bruto: R$ %.2f\n", resultado)
		fmt.Printf("Desconto do INSS: R$ %.2f\n", INSS)
		fmt.Printf("Salário líquido: R$ %.2f\n", resultado-INSS)
	} else if resultado >= 1518.01 && resultado <= 2793.88 {
		INSS = resultado * 0.090
		fmt.Printf("Salário bruto: R$ %.2f\n", resultado)
		fmt.Printf("Desconto do INSS: R$ %.2f\n", INSS)
		fmt.Printf("Salário líquido: R$ %.2f\n", resultado-INSS)
	} else if resultado >= 2793.89 && resultado <= 4190.83 {
		INSS = resultado * 0.120
		fmt.Printf("Salário bruto: R$ %.2f\n", resultado)
		fmt.Printf("Desconto do INSS: R$ %.2f\n", INSS)
		fmt.Printf("Salário líquido: R$ %.2f\n", resultado-INSS)
	} else if resultado >= 4190.84 && resultado <= 8157.41 {
		INSS = resultado * 0.140
		fmt.Printf("Salário bruto: R$ %.2f\n", resultado)
		fmt.Printf("Desconto do INSS: R$ %.2f\n", INSS)
		fmt.Printf("Salário líquido: R$ %.2f\n", resultado-INSS)
	} else {
		fmt.Println("Salário fora da faixa de cálculo do INSS. Consulte a tabela atualizada")
	}
}
