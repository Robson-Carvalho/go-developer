package main

import "fmt"

func main() {
    var x int = 42
    var ptrX *int = &x

    fmt.Printf(`
        Entendendo como ponteiros funciona
---------------------------------------
| Endereço   | Memória      |  Rótulo |
---------------------------------------
| %x | %v           | x       |
---------------------------------------
| %x | %v | ptrX    |
`, &x, x, &ptrX, ptrX)
}
