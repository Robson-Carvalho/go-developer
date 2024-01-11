package main

import "fmt"

func sayHello(){
    fmt.Println("Hello!👋")
}

func addTwoNumbers(x int, y int){
    total := 0
    total = x + y
    fmt.Println(total)
}


func addTwoNumbersWithReturn(x int, y int) int {
    total := 0
    total = x + y
    return total
}

func rectangleWithSigleReturn(l int, b int) (area int) {
    var parameter int
    parameter = 2 * (l+b)
    fmt.Println("Perímetro:", parameter)

    area = l * b
    return area
}

func rectangleWithMutipleReturns(l int, b int) (area int, parameter int) {
    parameter = 2 * (l+b)
    fmt.Println("Perímetro:", parameter)

    area = l * b
    return area, parameter
}

func update(a *int, t *string){
    fmt.Println("O local onde o ponteiro a está apontando na memória", a)
    fmt.Println("O local onde o ponteiro t está apontando na memória", t)
    fmt.Println("O endereço de memória do ponteiro a", &a)
    fmt.Println("O endereço de memória do ponteiro t", &t)

    *a = *a + 1
    *t = *t + " Carvalho"
}


func main() {
    sayHello()

    addTwoNumbers(20, 30)
    sum := addTwoNumbersWithReturn(30, 40)
    fmt.Println(sum)

    fmt.Println("Área:", rectangleWithSigleReturn(2, 3))
    var area, parameter int
    area, parameter = rectangleWithMutipleReturns(10, 20)
    fmt.Println("Área:", area, "\nPerímetro:", parameter)

    var age = 19
    var name = "Robson"
    fmt.Println("O endereço de memória da variável age", &age)
    fmt.Println("O endereço de memória da variável name", &name)
    fmt.Println("Before:", name, age )
    update(&age, &name)
    fmt.Println("After:", name, age )
}

