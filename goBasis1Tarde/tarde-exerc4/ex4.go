package main

import "fmt"

func main() {
	var counter int
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	//Imprime a idade de Benjamim
	fmt.Println("A idade de Benjamim é: ", employees["Benjamin"])

	//Adiciona um novo funcionário à lista, chamado Federico, que tem 25 anos.
	employees["Frederico"] = 21
	fmt.Println(employees)
	
	//Verifica quantos funcionários têm mais de 21 anos.
	for key := range employees {
		// sfmt.Println("Key:", key, "=>", "Element:", element)
			if employees[key] > 21 {
				counter += 1
			}
	}
	fmt.Println(counter, "funcioários tem mais de 21 anos")
	
}