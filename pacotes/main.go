package main

import (
	"fmt"
	"pacotes/math"
)

func main() {
	// Utilizando a função Soma do pacote matematica
	resultado := math.Soma(3, 5)
	fmt.Println("Resultado da Soma:", resultado)

    // pela função "subtração" começa com letra minúscula, não podemos usar em pacotes externos.
    /*resultado2 := math.subtracao(3, 5)
	fmt.Println("Resultado da Soma:", resultado2)
    */
}
