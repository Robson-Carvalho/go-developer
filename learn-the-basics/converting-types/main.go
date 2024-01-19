package main

import "fmt"

func main() {
  /* type inference - o compilador entender o tipo
  que está a variável está recebendo e não precisamos
  escrever de fato o tipo.
  */
  idade := 3

  fmt.Println(idade)

  /*Type casting - conversão de forma explícita de
  um tipo para outro compatível.
  */
  var ano_semestre float64 = 2024.1
  var ano int = int(ano_semestre)
  fmt.Println(ano)
}
