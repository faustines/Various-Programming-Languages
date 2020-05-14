/*
ECS 140A Programming Hw3
Faustine Yiu 913062973
Source: ECS 140A Prolog Lecture, StackOverflow
*/

isUnion([], Set, Set).

isUnion([A|Set1], Set2, Union) :-
    member(A, Set2),!,
    isUnion(Set1, Set2, Union).

isUnion([A|Set1], Set2, [A|Union]) :-
    isUnion(Set1, Set2, Union).

isIntersection([],_,[]).

isIntersection([Set1|A], Set2, [Set1|Intersection]) :-
    member(Set1, Set2), !,
    isIntersection(A, Set2, Intersection).

isIntersection([_|A], Set2, Intersection) :-
    isIntersection(A, Set2, Intersection).

isEqual(Set1, Set2) :-
    isIntersection(Set1, Set2, Intersection),
    Intersection = Set1,
    check(Set1, Set2).

check([], []).

check([_|Set1], [_|Set2]) :-
    check(Set1, Set2).

powerSet([],[[]]).

powerSet([A|B],Set)  :-
	powerSet(B,Set1),
	mergeeach([A],Set1,Set2),
	mergeset(Set1,Set2,Set).

mergeeach(X,[A0|[]],[A1]):-
	mergeset(X,A0,A1).

mergeeach(X,[A0|B0],[A1|B1]):-
	mergeset(X,A0,A1),
	mergeeach(X,B0,B1).

mergeset([], Set, Set).

mergeset([X|Set1],Set2,[X|Set]):-
      mergeset(Set1,Set2,Set).
