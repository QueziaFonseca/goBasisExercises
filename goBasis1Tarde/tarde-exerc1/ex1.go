package main

import "fmt"

func main ()  {
	palavra:= "Diego"
	fmt.Println(len(palavra))

	// for letra := range palavra {
	for _, letra := range palavra {

		// fmt.Printf("ï: %v, letra: %c\n", i, letra )
		fmt.Printf("%c\n", letra)

	}
}

