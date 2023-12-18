package main

import (
	"fmt"
	"strings"
)

type glc struct {
	notTerm     []string
	alphabet    []string
	productions map[string][]string
	root        string
}

// type production struct {
// 	name       string
// 	transition []string
// }

func main() {
	// obtém entrada digitada pelo usuário
	var inputs string
	fmt.Print("Digite uma entrada: ")
	fmt.Scanf("%v\n", &inputs)

	// converte entrada em uma lista (cada símbolo do alfabeto -> 1 valor da lista)
	sliceInputs := strings.Split(inputs, "")

	grammar := glc{
		notTerm:  []string{"S", "A"},
		alphabet: []string{"a", "b"},
		productions: map[string][]string{
			"S": {"AA", "AS", "b"},
			"A": {"SA", "AS", "a"},
		},
		root: "S",
	}

	grammar.cyk(sliceInputs)
}

func (g *glc) cyk(input []string) {

}
