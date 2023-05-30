package main

import "fmt"

func main() {
	var idade uint
	var temEmprego bool
	var tempoDeAtividade uint
	var salario float64

	idade = 23

	temEmprego = true
	tempoDeAtividade = 2
	salario = 200.00

		// ter mais 22 anos + empregado + + 1 ano de atividade
		if idade >22 && temEmprego == true &&  tempoDeAtividade > 1 {
			fmt.Println("Você tem direito ao emprestimo")

			//sem juros para salario superior a US$ 100.000
			if salario > 100.00 {
				fmt.Println("Seu emprestimo será sem juros")
			}
		} else {
			fmt.Println("Você não tem direito ao emprestimo")
		}
	}
/*
	// ter mais 22 anos
	if idade <22 {
		fmt.Println("Você precisa ter pelo menos 22 anos")
	// empregado
	} else if  temEmprego == false {
		fmt.Println("Você precisa estar empregado")

	//+ 1 ano de atividade
	}else if tempoDeAtividade > 1 {
		fmt.Println("Você precisa ter pelo menos 22 anos")
	//sem juros para salario superior a US$ 100.000
		if salario > 100.00 {		
			fmt.Println("Você tem direito ao emprestimo sem juros")
		}
	}
*/