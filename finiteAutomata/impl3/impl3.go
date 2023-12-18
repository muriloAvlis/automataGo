package main

import (
	"fmt"
)

// definição do autômato
type automata struct {
	alphabet     []string         // alfabeto de símbolos de entrada
	exitAlphabet []string         // alfabeto de saída do transdutor
	states       map[string]state // conj. finito de estados possíveis do autômato
	initState    state            // estado inicial
	finalStates  map[string]state // conj. de estados finais
}

type state struct {
	name       string            // nome do estado
	transFunc  map[string]string // função de transição do estado
	transdFunc map[string]string // função de transdução
}

func main() {
	// obtém entrada digitada pelo usuário
	inputs := []string{"100", "25", "25", "25", "25", "100", "50", "50", "100", "100", "25", "50", "25", "50", "25", "25", "100"}
	fmt.Printf("Entrada: %v\n", inputs)

	// declaração dos estados do autômato
	q0 := state{
		name:       "q0",
		transFunc:  map[string]string{"100": "q0", "25": "q1", "50": "q2"},
		transdFunc: map[string]string{"100": "1", "25": "0", "50": "0"},
	}

	q1 := state{
		name:       "q1",
		transFunc:  map[string]string{"100": "q1", "25": "q2", "50": "q3"},
		transdFunc: map[string]string{"100": "1", "25": "0", "50": "0"},
	}

	q2 := state{
		name:       "q2",
		transFunc:  map[string]string{"100": "q1", "25": "q3", "50": "q0"},
		transdFunc: map[string]string{"100": "1", "25": "0", "50": "1"},
	}

	q3 := state{
		name:       "q3",
		transFunc:  map[string]string{"100": "q3", "25": "q0", "50": "q1"},
		transdFunc: map[string]string{"100": "1", "25": "1", "50": "1"},
	}

	// declaração do automato
	automata := automata{
		alphabet:     []string{"25", "50", "100"},
		exitAlphabet: []string{"0", "1"},
		states:       map[string]state{"q0": q0, "q1": q1, "q2": q2, "q3": q3},
		initState:    q0,
		finalStates:  map[string]state{"q0": q0, "q1": q1, "q2": q2, "q3": q3},
	}

	// aplica a lógica de transição para a entrada
	automata.transLogic(inputs)
}

// função que implementa a lógica de transição entre os estados
func (a *automata) transLogic(inputs []string) {
	var currentState state // estado atual
	var nextState string   // próximo estado
	var output []string

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
		// guarda a saída de acordo com a transição que foi feita
		output = append(output, currentState.transdFunc[input])
	}

	fmt.Printf("Saída: %v\n", output)
}
