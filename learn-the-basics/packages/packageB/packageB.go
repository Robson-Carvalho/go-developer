// pacoteB.go
package pacoteB

import (
	"fmt"
	"packages/packageA"
)

func CumprimentoPersonalizado(nome string) {
    pacoteA.Saudacao(nome)
    fmt.Println("Bem-vindo ao pacoteB!")
}
