package main

import (
	"fmt"

	"github.com/robson-carvalho/golang-studies/praticando/math"
)

var nome string

func main() {
    nome = "Robson"

    // Inferência de tipo - Não dizemos o tipo, mas pelo dado recebido um tipo é determinado.
    idade := 19

    fmt.Printf("Eu me chamo %v e tenho %v anos\n", nome, idade)

    soma := math.Soma(1, 2)

    hello := math.Hello

    fmt.Printf("%T - %v \n", hello, hello)

    fmt.Printf("%T - %v \n", soma, soma)
}

