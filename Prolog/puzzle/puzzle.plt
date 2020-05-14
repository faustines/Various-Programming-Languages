:- initialization main.

main :-
    consult(["puzzle.pl"]),
    (show_coverage(run_tests) ; true),
    halt.

:- begin_tests(puzzle).

% It is possible to move farmer, wolf, goat and cabbage from left bank to right bank.
test(puzzle1, [true, nondet]) :-
    solve(left,left,left,left,right,right,right,right).
% It is NOT possible to start with farmer and cabbage on the left bank
% and the wolf and goat on the right bank,
% and then end with farmer, wolf, goat and cabbage on the right bank.
% (The wolf will eat the goat when left on the right bank without the farmer.)
test(puzzle2, [fail, nondet]) :-
    solve(left,right,right,left,right,right,right,right).
% It is NOT possible to start with farmer, wolf, goat and cabbage on the left bank,
% and end up with the farmer and wolf on the right bank, and the goat and cabbage
% on the left bank.
% (The goat will eat the cabbage when left on the left bank without the farmer.)
test(puzzle3, [fail, nondet]) :-
    solve(left,left,left,left,right,right,left,left).

:- end_tests(puzzle).
