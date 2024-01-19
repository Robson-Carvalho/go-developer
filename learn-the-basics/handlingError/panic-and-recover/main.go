package main

import (
	"fmt"
)

func main() {
	// Exemplo 1: Panic sem recover
	// Neste caso, o programa entra em pânico e encerra abruptamente.
	// A mensagem associada ao pânico é impressa.
	// Uncomment para testar
	//panicExample1()

	// Exemplo 2: Panic com recover
	// Neste caso, o recover() dentro da função de recuperação é capaz de capturar
	// o pânico, imprimir a mensagem associada e continuar a execução.
	// Uncomment para testar
	//panicExample2()

	// Exemplo 3: Uso de panic e recover em uma função que retorna erro
	// Aqui, estamos usando panic para tratar um erro crítico e recover para
	// recuperar-se dele na função de recuperação.
	// Uncomment para testar
	//errorHandlingExample()

	// Exemplo 4: Uso de defer para recuperar de panics em funções chamadas
	// por outras funções
	// Aqui, uma função chama outra função que entra em pânico. O defer em
	// recoverFromPanic() é capaz de recuperar o pânico da função chamada.
	// Uncomment para testar
	deferExample()
}

func panicExample1() {
	fmt.Println("Exemplo 1: Panic sem recover")
	panic("Algo deu errado!")
	fmt.Println("Esta linha não será alcançada.")
}

func panicExample2() {
	fmt.Println("Exemplo 2: Panic com recover")
	defer recoverFromPanic()
	panic("Algo deu errado, mas conseguimos recuperar!")
}

func errorHandlingExample() {
	fmt.Println("Exemplo 3: Uso de panic e recover em uma função que retorna erro")
	result, err := divideWithErrorHandling(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", result)
	}
}

func divideWithErrorHandling(a, b int) (int, error) {
	if b == 0 {
		panic("Não é possível dividir por zero!")
	}
	return a / b, nil
}

func deferExample() {
	fmt.Println("Exemplo 4: Uso de defer para recuperar de panics em funções chamadas por outras funções")
	defer recoverFromPanic()
	callFunctionThatPanics()
}

func callFunctionThatPanics() {
	fmt.Println("Chamando função que entra em pânico...")
    /* Quando entra em panic, o programa para e volta para as
    funções passadas, buscando um recover, caso não ache, temos
    um erro e parada do programa.
    */
	panic("Pânico na função chamada!")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado de um pânico:", r)
	}
}
