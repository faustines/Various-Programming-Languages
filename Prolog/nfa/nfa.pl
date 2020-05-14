/*
ECS 140A Programming Hw3
Faustine Yiu 913062973
Source: ECS 140A Prolog Lecture, StackOverflow, cpp.edu
*/

reachable(StartState,FinalState,Input) :-
  StartState == FinalState -> true;
  transit(StartState,FinalState,Input).

transit(S,F,[A|B]) :-
        transition(S,A,[Y|_]),
        nl,
        transit(Y,F,B).

transit(S,F,[A|B]) :-
        transition(S,A,[_|Y]),
        nl,
        transit(Y,F,B).

transit(S,F,[]) :-
        S == F -> true;
        false.

transit([S|_],F,[]) :-
        S == F -> true;
        false.
