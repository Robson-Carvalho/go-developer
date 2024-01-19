package main

/*
Em go, geralmente, um erro é representado por uma interface
chamada 'error', que possui um único método 'Error()' que
retorna uma string descrevendo o erro. Exemplo de uso de erro
em Go: A função dividir retorna dois parãmetros e se o segundo
não for nulo, entramos no if e apresentamos o erro.
*/

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}


func main() {
    op1, err := divide(2, 2)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(op1)

    op2, err := divide(2, 0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(op2)

    op3, err := divide(3, 0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(op3)

}
