package main

import (
	"fmt"
	"strings"
)

// definição do tipo autômato com pilha
type automata struct {
	alphabet    []string         // alfabeto de símbolos de entrada
	states      map[string]state // conj. finito de estados possíveis do autômato
	initState   state            // estado inicial
	finalStates map[string]state // conj. de estados finais
}

// definição do tipo estado
type state struct {
	name      string                  // nome do estado
	transFunc map[[2]string][2]string // função de transição do estado (a, Z) -> (Qn, XZ)
}

func main() {
	// obtém entrada digitada pelo usuário
	var inputs string
	fmt.Print("Digite uma entrada: ")
	fmt.Scanf("%v\n", &inputs)

	// converte entrada em uma lista (cada símbolo do alfabeto -> 1 valor da lista)
	sliceInputs := strings.Split(inputs, "")

	// declaração dos estados do autômato
	q0 := state{
		name: "q0",
		transFunc: map[[2]string][2]string{
			{"a", "Z"}: {"q1", "XZ"},
			{"c", "Z"}: {"q3", "Z"}},
	}

	q1 := state{
		name: "q1",
		transFunc: map[[2]string][2]string{
			{"a", "X"}: {"q1", "XX"},
			{"b", "X"}: {"q2", ""},
		},
	}

	q2 := state{
		name: "q2",
		transFunc: map[[2]string][2]string{
			{"b", "X"}: {"q2", ""},
			{"c", "Z"}: {"q3", "Z"},
		},
	}

	q3 := state{
		name: "q3",
		transFunc: map[[2]string][2]string{
			{"c", "Z"}: {"q3", "Z"},
		},
	}

	// declaração do automato
	pushDownAutomata := automata{
		alphabet:    []string{"a", "b", "c"},
		states:      map[string]state{"q0": q0, "q1": q1, "q2": q2, "q3": q3},
		initState:   q0,
		finalStates: map[string]state{"q3": q3},
	}

	pushDownAutomata.transLogic(sliceInputs)
}

// função que implementa a lógica de transição entre os estados
func (a *automata) transLogic(inputs []string) {
	if len(inputs) == 0 { // caso a entrada seja vazia, ela é rejeitada
		fmt.Printf("Entrada %v rejeitada!\n", inputs)
	} else {
		var currentState state // estado atual
		var nextState string   // próximo estado
		stack := []string{"Z"} // conteúdo da pilha

		started := false               // váriavel de controle
		for _, input := range inputs { // iteração sobre cada símbolo da entrada
			if started != true { // se ainda não inicializou, começar pelo estado inicial
				currentState = a.initState
				started = true
			} else { // se já inicializou, obtém o estado corrente
				currentState = a.states[nextState]
			}

			// obtém o próximo estado, aplicando a função de transição do estado corrente (a, Z) -> (Qn, XZ)
			nextState = currentState.transFunc[[2]string{input, stack[0]}][0]
			if nextState == "" { // transição inválida
				break
			}
			toStack := strings.Split(currentState.transFunc[[2]string{input, stack[0]}][1], "") // o que vai ser adicionado a pilha
			stack = stack[1:]                                                                   // remove o item no topo da pilha
			stack = append(toStack, stack...)
		}

		currentState = a.states[nextState] // obtém o ultimo estado da iteração

		// verifica se o estado obtido está contido no conjunto de estados finais
		_, ok := a.finalStates[currentState.name]

		// mostra mensagem de aceitação caso esteja no conjunto de estados finais ou rejeitação caso não
		if ok {
			fmt.Printf("Entrada %v aceita!\n", inputs)
			fmt.Printf("Conteúdo da pilha: %v\n", stack)
		} else {
			fmt.Printf("Entrada %v rejeitada!\n", inputs)
			fmt.Printf("Conteúdo da pilha: %v\n", stack)
		}
	}
}
