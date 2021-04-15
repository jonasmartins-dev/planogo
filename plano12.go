package main

import (
	"fmt"
)

func main() {

	var (
		peso   float64 = 75.5
		idade  int     = 18
		altura float64 = 1.71
	)
	
	fmt.Println("Olá")
	fmt.Println("Sua idade é igual a: ", idade)
	fmt.Printf("Você apresentou um peso de %f quilos e uma altura de %f metros", peso, altura)

}
