package main

import (
	"fmt"
	"strings"
)

// definição do autômato
type automata struct {
	alphabet    []string         // alfabeto de símbolos de entrada
	states      map[string]state // conj. finito de estados possíveis do autômato
	initState   state            // estado inicial
	finalStates map[string]state // conj. de estados finais
}

type state struct {
	name      string            // nome do estado
	transFunc map[string]string // função de transição do estado
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
		name:      "q0",
		transFunc: map[string]string{"a": "q1", "b": "q2"},
	}

	q1 := state{
		name:      "q1",
		transFunc: map[string]string{"a": "q0", "b": "q1"},
	}

	q2 := state{
		name:      "q2",
		transFunc: map[string]string{"a": "q1", "b": "q2"},
	}

	// declaração do automato
	automata := automata{
		alphabet:    []string{"a", "b"},
		states:      map[string]state{"q0": q0, "q1": q1, "q2": q2},
		initState:   q0,
		finalStates: map[string]state{"q2": q2},
	}

	// aplica a lógica de transição para a entrada
	automata.transLogic(sliceInputs)
}

// função que implementa a lógica de transição entre os estados
func (a *automata) transLogic(inputs []string) {
	var currentState state // estado atual
	var nextState string   // próximo estado

	started := false // váriavel de controle

	for _, input := range inputs { // iteração sobre cada símbolo da entrada
		if started != true { // se ainda não inicializou, começar pelo estado inicial
			currentState = a.initState
			started = true
		} else { // se já inicializou, obtém o estado corrente
			currentState = a.states[nextState]
		}

		// obtém o próximo estado, aplicando a função de transição do estado corrente
		nextState = currentState.transFunc[input]
	}

	currentState = a.states[nextState] // obtém o ultimo estado da iteração

	// verifica se o estado obtido está contido no conjunto de estados finais
	_, ok := a.finalStates[currentState.name]

	// mostra mensagem de aceitação caso esteja no conjunto de estados finais ou rejeitação caso não
	if ok {
		fmt.Printf("Entrada %v aceita!\n", inputs)
	} else {
		fmt.Printf("Entrada %v rejeitada!\n", inputs)
	}
}
