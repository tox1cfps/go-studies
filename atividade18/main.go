package main

import (
	"fmt"
	"strings"
)

func main() {

	var senhas []string

	for {
		var senha string
		fmt.Print("Digite sua senha: ")
		fmt.Scan(&senha)

		simbolos := "!@#$%^&*()-_=+[]{};:,.<>?/|~`"

		var temMaiuscula bool
		var temMinuscula bool
		var temNumero bool
		var temEspecial bool

		for _, caractere := range senha {
			if caractere >= 'A' && caractere <= 'Z' {
				temMaiuscula = true
			} else if caractere >= 'a' && caractere <= 'z' {
				temMinuscula = true
			} else if caractere >= '0' && caractere <= '9' {
				temNumero = true
			} else if strings.ContainsRune(simbolos, caractere) {
				temEspecial = true
			}
		}

		if len(senha) < 8 {
			fmt.Println("A senha precisa ter no minimo 8 caracteres!")
		} else if !temMinuscula {
			fmt.Println("A senha precisa conter no minimo uma letra minuscula")
		} else if !temMaiuscula {
			fmt.Println("A senha precisa conter no minimo uma letra maiuscula")
		} else if !temNumero {
			fmt.Println("A senha precisa conter no minimo um numero")
		} else if !temEspecial {
			fmt.Println("A senha precisa conter no minimo um caractere especial")
		} else {
			senhas = append(senhas, senha)
			fmt.Println("Senha criada com sucesso")
			break
		}
	}
}
