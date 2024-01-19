package main

import "fmt"


func main(){
    fmt.Println("Números pares de 0 a 50")

    for i := 0; i <= 50; i ++ {
        if (i % 2 == 0){
            fmt.Println(i)
        }
    }

    fmt.Println("A soma de 1 até 1000")
    var soma int = 0

    for e := 0; e <= 1000; e++ {
        soma += e
    }

    fmt.Printf("A soma é: %v\n", soma)

    //While em Go
    condicao := 1
    for condicao < 5 {
        condicao *= 2
    }

    fmt.Println(condicao)

    //For-each
    strings := []string{"hello", "world", "go", "lang", "backend"}
    for i, s := range strings {
	    fmt.Println("Index: ", i, "- Valor:", s)
    }
}
