package main

import "fmt"

func main() {
    // Declarar e inicializar um map de strings para inteiros
    idadePorNome:= make(map[string]int)

    // Adicionar elementos ao map
    idadePorNome["Alice"] = 25
    idadePorNome["Bob"] = 30

    // Recuperar e imprimir valores do map
    fmt.Println("Idade de Alice:", idadePorNome["Alice"]) // Saída: Idade de Alice: 25
    fmt.Println("Idade de Bob:", idadePorNome["Bob"])     // Saída: Idade de Bob: 30

    idade, existe := idadePorNome["Charlie"]

    if existe {
      fmt.Println("Idade de Charlie:", idade)
    } else {
      fmt.Println("Charlie não encontrado")
    }



    // Declarar e inicializar um map com string como chave e slice de inteiros como valor
    dadosPorCategoria := make(map[string][]int)

    // Adicionar elementos ao map
    dadosPorCategoria["números pares"] = []int{2, 4, 6, 8, 10}
    dadosPorCategoria["números ímpares"] = []int{1, 3, 5, 7, 9}

    // Acessar e imprimir valores do map
    fmt.Println("Números pares:", dadosPorCategoria["números pares"])     // Saída: Números pares: [2 4 6 8 10]
    fmt.Println("Números ímpares:", dadosPorCategoria["números ímpares"]) // Saída: Números ímpares: [1 3 5 7 9]

    dadosPorCategoria["números pares"] = append(dadosPorCategoria["números pares"], 2024)

    // Acessar e imprimir valores do map
    fmt.Println("Números pares:", dadosPorCategoria["números pares"])     // Saída: Números pares: [2 4 6 8 10]
    fmt.Println("Números ímpares:", dadosPorCategoria["números ímpares"]) // Saída: Números ímpares: [1 3 5 7 9
}
