:- initialization main.

main :-
    consult(['transitions.pl', 'nfa.pl']),
    (show_coverage(run_tests) ; true),
    halt.

:- begin_tests(nfa).

test(nfa1, [nondet]) :- reachable(0, 2, [a]).

test(nfa2, [nondet]) :-  reachable(0, 2, [b]).

test(nfa3, [nondet]) :-  reachable(0, 1, [a, b, a]).

test(nfa4, [fail]) :- reachable(0, 1, [a, b, a, b]).

test(nfa5, [nondet]) :- reachable(0, 2, [a, b, a]).

:- end_tests(nfa).
