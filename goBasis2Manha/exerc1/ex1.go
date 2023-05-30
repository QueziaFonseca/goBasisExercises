package main

// import "fmt"
var salario float64
var desconto float64
var imposto float64

func calImposto(salario, desconto, imposto float64) (float64,error) {
	if salario == 50000.00 {
		desconto = salario*0.17
		imposto = salario - desconto
		return imposto, nil
	}

	if salario == 150000.00 {
		desconto = salario*0.10
		imposto = salario - desconto
		return imposto, nil
	}
	return error.new(string)
		// ("Não é possível calcular o imposto sobre o seu salário")
}

func main ()  {
	
}

