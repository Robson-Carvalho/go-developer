package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main(){
    res, err := http.Get("http://google.com.br")
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Println(res.Body)



    soma, err := somaF(1, 29)
    if err != nil{
        log.Fatal(err.Error())
    }

    fmt.Printf("o valor da soma é: %v\n", soma)


    // Uma forma de não tratar o erro
    soma2, _ := somaF(2,30)

    fmt.Printf("o valor da soma é: %v\n", soma2)
}

func somaF(x int, y int) (int, error){
    res := x + y

    if res > 10 {
        return 0, errors.New("Total maior que 10")
    }

    return res, nil
}
