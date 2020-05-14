/*
ECS 140a Hw2
Faustine Yiu 913062973
*/
package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO Write the Reachable function,
	// return true if the nfa accepts the input,
	// return false if it doesn't reach to the final state with that input
	var next []state

	//if the input is empty, but start state remains at final state return true, else return false
	if len(input) == 0{
		if start == final {
			return true
		}else{
			return false
		}
	}

	next = transitions(start,input[0]) //first rune transition stored in next

	//loop through rune transitions combinations and filter out the deadends with != 0
	//set next to the next rune transition combinations
	//if it gets to a deadend, check if the alphabet is a (for langTransitions a loops), if it is not a, return false
	for i:= 1; i < len(input); i++ {
		for j:= 0; j < len(next); j++ {
			if (len(transitions(next[j],input[i])) != 0) {
				next = transitions(next[j],input[i])
			}else{
				if input[i-1] != 'a' && input[i] != 'a'{
					return false
				}
			}
		}
	}

	//loop through next to see if any maps to final
	for k:= 0; k < len(next); k++{
			if next[k] == final {
				return true
			}
	}
	return false
}
