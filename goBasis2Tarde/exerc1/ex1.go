package main

import "fmt"

// Nome: [Nome do aluno]
// Sobrenome: [Sobrenome do aluno]
// RG: [RG do aluno]
// Data de admissão: [Data de admissão do aluno]
var nome string
var sobrenome string
var RG string
var dataDeAdmissão string

func cadastro (nome, sobrenome, RG, dataDeAdmissão string)   (string string string string) {
	return nome, sobrenome, RG, dataDeAdmissão
}

func main() {
	fmt.Println(cadastro(nome, sobrenome, RG, dataDeAdmissão))
}