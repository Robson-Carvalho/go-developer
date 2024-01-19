/* Em Go (ou Golang), a palavra-chave range é utilizada em loops
para iterar sobre elementos de estruturas de dados como slices,
arrays, maps, strings, e canais. O range retorna dois valores: o
índice (ou chave) e o valor do elemento na iteração atual. A
forma básica de usar o range é em um loop for.
*/

package main

import "fmt"

func main(){
    var nomes = []string{"Robson", "Rodrigo", "Roque", "Rosimeire"}

    for i, nome := range nomes {
        fmt.Println(i, nome)
    }

    texto := "日本語"

    for i, ch := range texto {
        fmt.Printf("%#U starts at byte position %d\n", ch, i)
    }


    for i := 0; i < len(texto); i++ {
	    fmt.Printf("%x \n", texto[i])
    }

    dados := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
    }

    for chave, valor := range dados {
        fmt.Println(chave, valor)
    }
}
