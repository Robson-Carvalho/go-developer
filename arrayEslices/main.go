package main

import "fmt"

func main() {
  // Array- estrura de dados estática que armazena uma coleção de dados do mesmo tipo
  //primeira forma de declarar
  fmt.Println("\nArrays")
  var listaDeNumeros1 [4]int

  for i := 0; i < len(listaDeNumeros1); i++ {
    fmt.Println(listaDeNumeros1[i])
  }

  // segunda forma de declarar
  listaDeNumeros2 := [3]int{2}

  for i := 0; i < len(listaDeNumeros2); i++ {
    fmt.Println(listaDeNumeros2[i])
  }

  // Slices - Estrutura de dados dinâmica que representa parte de um array
  // declaração básica de um slice - declaramos como um array, só não colocamos o tamanho.
  fmt.Println("\nSlices")
  var meuSlice []int
  contador := 10
  for i := 0; i < contador; i++ {
    meuSlice = append(meuSlice, i)
  }

  for i := 0; i < contador; i++ {
    fmt.Println(meuSlice[i])
  }
}
