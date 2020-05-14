/*
ECS 140A Programming Hw5
Faustine Yiu 913062973
Source: https://golang.org, ECS 140A Go Lecture, StackOverflow
*/

package nfa
import "sync"

// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, act rune) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO
	returns := make(chan bool)
	var wg sync.WaitGroup
	reachable_helper(input,start,final,0,returns,&wg,transitions)
	go func(){
	    wg.Wait()
	    close(returns)
	}()
	return <- returns
}

func reachable_helper (input []rune,start,final state,num int,returns chan bool,wg *sync.WaitGroup,transitions TransitionFunction){
     wg.Add(1)
	    go func(){
	        trans(input,start,final,num,returns,wg,transitions)
	        wg.Done()
	    }()
}

func trans(input []rune,start,final state,num int,returns chan bool,wg *sync.WaitGroup,transitions TransitionFunction){
    if (num == (len(input))){
        if (start == final){
            returns <- true
				}
    }else{
        states:= transitions(start,input[num])
        for _,possible := range states{
                reachable_helper(input,possible,final,num+1,returns,wg,transitions)
        }
    }
}
