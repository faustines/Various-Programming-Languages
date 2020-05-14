/*
ECS 140A Programming Hw3
Faustine Yiu 913062973
Source: ECS 140A Prolog Lecture, StackOverflow
*/

solve(F1, W1, G1, C1, F2, W2, G2, C2) :-
not(unsafe(F1, W1, G1, C1)),
not(unsafe(F2, W2, G2, C2)),
move(state(F1, W1, G1, C1),state(F2, W2, G2, C2)), !.

opp(right,left).
opp(left,right).

% wolf eats goat
unsafe(right,left,left,_).
unsafe(left,right,right,_).

% goat eats cabbage
unsafe(right,_,left,left).
unsafe(left,_,right,right).

% Goal state
move(state(right,right,right,right),_) :- nl.

% Take farmer and wolf
move(state(F,F,G1,C1),L) :-
 opp(F,F2),
 opp(F,W2),
 G2 = G1,
 C2 = C1,
 not(unsafe(F2,W2,G2,C2)),
 not(member(state(F2,W2,G2,C2),L)),
 move(state(F2,W2,G2,C2),[state(F,F,G1,C1)|L]).

% Take farmer and goat
move(state(F,W1,F,C1),L) :-
 opp(F,F2),
 opp(F,G2),
 W2 = W1,
 C2 = C1,
 not(unsafe(F2,W2,G2,C2)),
 not(member(state(F2,W2,G2,C2),L)),
 move(state(F2,W2,G2,C2),[state(F,W1,F,C1)|L]).

% Take farmer and cabbage
move(state(F,W1,G1,F),L) :-
 opp(F,F2),
 opp(F,C2),
 G2 = G1,
 W2 = W1,
 not(unsafe(F2,W2,G2,C2)),
 not(member(state(F2,W2,G2,C2),L)),
 move(state(F2,W2,G2,C2),[state(F,W1,G1,F)|L]).

% Take farmer and nobody else
move(state(F,W1,G1,C1),L) :-
 opp(F,F2),
 C2 = C1,
 G2 = G1,
 W2 = W1,
 not(unsafe(F2,W2,G2,C2)),
 not(member(state(F2,W2,G2,C2),L)),
 move(state(F2,W2,G2,C2),[state(F,W1,G1,C1)|L]).
