package main

import "fmt"
/* struct contendo
Nome: [Nome do aluno]
Sobrenome: [Sobrenome do aluno]
RG: [RG do aluno]
Data de admissão: [Data de admissão do aluno]*/

type aluno struct {
	Nome string
	Sobrenome string
	RG string
	DataDeAdmissão string
}


func (a aluno) detalhe() {
	fmt.Println("Nome do aluno:", a.Nome)
	fmt.Println("Sobrenome do aluno:", a.Sobrenome)
	fmt.Println("RG do aluno:", a.RG)
	fmt.Println("Data de Admissão do aluno:", a.DataDeAdmissão)
}

func main() {
	aluno1 := aluno{
		Nome: "Diego",
		Sobrenome: "Limas",
		RG: "123456",
		DataDeAdmissão: "01/01/2023",
	}
	aluno1.detalhe()
}