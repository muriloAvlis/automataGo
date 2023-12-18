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

func (g *glc) cyk(inputs []string) {
	V1C := []string{}
	VLC := []string{}
	for _, input := range inputs {
		for k, v := range g.productions {
			for _, prod := range v {
				if prod == input {
					V1C = append(V1C, k)
				}
			}
		}
	}

	for i, c := range V1C {
		if i < len(V1C)-1 {
			var temp string
			word := c + V1C[i+1]
			for k, v := range g.productions {
				for _, prod := range v {
					if prod == word {
						temp += k
					}
				}
			}
			VLC = append(VLC, temp)
		}
	}

	fmt.Println(VLC[:len(VLC)-3])
	fmt.Println(VLC[:len(VLC)-2])
	fmt.Println(VLC[:len(VLC)-1])
	fmt.Println(VLC)
	fmt.Println(V1C)

	test := strings.Join(VLC[:len(VLC)-3], "")
	testSplit := strings.Split(test, "")

	for _, v := range testSplit {
		if v == g.root {
			fmt.Println("Entrada aceita!")
			return
		}
	}

	fmt.Println("Entrada rejeitada!")
}
