package main

import "fmt"
var salario float64
var desconto float64
var imposto float64

func calImposto(salario float64) (float64) {
	if salario <= 50000.00 {
		return 0.0
		} else if salario <= 150000.00 {
		return salario * 0.17
		}
		return salario * 0.3
	}

func main ()  {
fmt.Println("imposto de até R$50.000: ", calImposto(50000))
fmt.Println("imposto de até R$150.000: ", calImposto(150000))
fmt.Println("imposto de até R$50.000: ", calImposto(150010))
}

