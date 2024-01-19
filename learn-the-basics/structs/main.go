package main

import "fmt"

type Produto struct {
  nome string
  valor float64
  quantidade int
}

type Pessoa struct {
  Nome  string
  Idade int
}


func (p Produto) ValorTotal() float64 {
  valorTotal := p.valor * float64(p.quantidade)

  return valorTotal
}

func main() {
  produtos := []Produto{}

  p1 := Produto{nome: "Queijo 500g", valor: 6.84, quantidade: 5}
  p2 := Produto{nome: "Suco de uva 1L", valor: 13.99, quantidade: 2}
  p3 := Produto{nome: "Pão", valor: 4.50, quantidade: 1}
  p4 := Produto{nome: "Presunto 500g", valor: 3.84, quantidade: 5}

  produtos = append(produtos, p1, p2, p3, p4)

  var valorTotal float64

  for i := 0; i < len(produtos); i++ {
    valorTotal += produtos[i].ValorTotal()
    fmt.Printf("O produto %s vale R$ %.2f e foi comprado %d, custando R$ %.2f", produtos[i].nome, produtos[i].valor, produtos[i].quantidade, produtos[i].ValorTotal())
    fmt.Println()
  }
  fmt.Printf("\nO valor total da compra foi R$ %.2f\n", valorTotal)


  pessoa := Pessoa{"Alice", 25}

  // Passagem por valor (cópia)
  copia := outraFuncao(pessoa)
  copia.Nome = "Charlie"
  fmt.Println(pessoa.Nome) // Saída: Alice

  // Passagem por referência
  referencia := outraFuncaoComPonteiro(&pessoa)
  referencia.Nome = "Eva"
  fmt.Println(pessoa.Nome) // Saída: Eva
}

func outraFuncao(p Pessoa) Pessoa {
    return p
}

func outraFuncaoComPonteiro(p *Pessoa) *Pessoa {
    return p
}


