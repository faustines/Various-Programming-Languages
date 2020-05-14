/*
ECS 140A Programming Hw3
Faustine Yiu 913062973
Source: ECS 140A Prolog Lecture, StackOverflow
*/

year_1953_1996_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    novel(Book,1953);
    novel(Book,1996).

period_1800_1900_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    novel(Book,Y),
    Y >= 1800,
    Y =< 1900.

lotr_fans(Fan) :-
    %% remove fail and add body/other cases for this predicate
    fan(Fan, B),
    member(the_lord_of_the_rings, B).

author_names(Author) :-
    %% remove fail and add body/other cases for this predicate
    fan(chandler, C),
    author(Author, A),
    any_member(C, A).

fans_names(Fan) :-
    %% remove fail and add body/other cases for this predicate
    author(brandon_sanderson, BS),
    fan(Fan, F),
    any_member(BS, F).

%% Find all elements of one List which are also members of another List
any_member([A|_], [A|_]) :- !.
any_member(S1, [_|T2]) :- any_member(S1, T2), !.
any_member([_|T1], S2) :- any_member(T1, S2), !.

mutual_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    fan(phoebe, P),
    fan(ross, R),
    member(Book, P),
    member(Book, R).

mutual_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    fan(monica, M),
    fan(ross, R),
    member(Book, M),
    member(Book, R).

mutual_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    fan(phoebe, P),
    fan(monica, M),
    member(Book, P),
    member(Book, M).
