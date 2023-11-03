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

// definição do estado do autômato
type state struct {
	name      string            // nome do estado
	transFunc map[string]string // função de transição do estado
}

func main() {
	// obtém entrada digitada pelo usuário
	inputs := `O computador é uma máquina capaz de variados tipos de tratamento automático de informações ou processamento de dados.
Entende-se por computador um sistema físico que realiza algum tipo de computação. Assumiu-se que os computadores pessoais e
laptops são ícones da era da informação. O primeiro computador eletromecânico foi construído por Konrad Zuse (1910–1995). Atualmente,
um microcomputador é também chamado computador pessoal ou ainda computador doméstico.`

	// converte entrada em uma lista (cada símbolo do alfabeto -> 1 valor da lista)
	sliceInputs := strings.Split(inputs, "")

	// declaração dos estados do autômato
	q0 := state{
		name:      "q0",
		transFunc: map[string]string{"c": "q1"},
	}

	q1 := state{
		name:      "q1",
		transFunc: map[string]string{"o": "q2"},
	}

	q2 := state{
		name:      "q2",
		transFunc: map[string]string{"m": "q3"},
	}

	q3 := state{
		name:      "q3",
		transFunc: map[string]string{"p": "q4"},
	}

	q4 := state{
		name:      "q4",
		transFunc: map[string]string{"u": "q5"},
	}

	q5 := state{
		name:      "q5",
		transFunc: map[string]string{"t": "q6"},
	}

	q6 := state{
		name:      "q6",
		transFunc: map[string]string{"a": "q7"},
	}

	q7 := state{
		name:      "q7",
		transFunc: map[string]string{"d": "q8"},
	}

	q8 := state{
		name:      "q8",
		transFunc: map[string]string{"o": "q9"},
	}

	q9 := state{
		name:      "q9",
		transFunc: map[string]string{"r": "q10"},
	}

	q10 := state{
		name:      "q10",
		transFunc: map[string]string{" ": "q0"},
	}

	// declaração do automato
	automata := automata{
		alphabet:    []string{"c", "o", "m", "p", "u", "t", "a", "d", "o", "r"},
		states:      map[string]state{"q0": q0, "q1": q1, "q2": q2, "q3": q3, "q4": q4, "q5": q5, "q6": q6, "q7": q7, "q8": q8, "q9": q9, "q10": q10},
		initState:   q0,
		finalStates: map[string]state{"q10": q10},
	}

	// aplica a lógica de transição para a entrada
	automata.transLogic(sliceInputs)
}

// função que implementa a lógica de transição entre os estados
func (a *automata) transLogic(inputs []string) {
	var currentState state // estado atual
	var nextState string   // próximo estado
	var idxs []int         // posições em que ocorrem a palavra computador

	started := false // váriavel de controle

	for idx, input := range inputs { // iteração sobre cada símbolo da entrada
		if started != true { // verifica inicialização e se começa com c
			currentState = a.initState
			started = true
		} else { // se já inicializou, obtém o estado corrente
			currentState = a.states[nextState]
		}
		// obtém o próximo estado, aplicando a função de transição do estado corrente
		nextState = currentState.transFunc[input]

		if currentState.name != "q10" && nextState == "" { // não forma a sequência computador e volta pro início
			nextState = "q0"
		} else if currentState.name == "q10" && nextState != "" { // palavra aceita no início/meio do texto
			if len(inputs[:idx]) > 10 { // caso a palavra não esteja no início
				if inputs[idx-11] == " " { // verifica se existe prefixo
					idxs = append(idxs, idx-10)
				}
			} else { // palavra no início do texto
				idxs = append(idxs, idx-10)
			}
		} else if (nextState == "q10") && (len(inputs)-1 == idx) { // computador como última palavra
			if inputs[idx-10] == " " { // verifica se existe prefixo antes de computador
				idxs = append(idxs, idx-9)
			}
		}
	}

	// mostra os índices obtidos
	fmt.Println(idxs)
}
