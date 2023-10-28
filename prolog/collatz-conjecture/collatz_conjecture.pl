collatz_steps(N, Steps) :- N > 0 -> collatz_steps(N, Steps, 0).
collatz_steps(N, Steps, Count) :- N =:= 1, Steps = Count.
collatz_steps(N, Steps, Count) :- N mod 2 =:= 0, M is N / 2, NewCount is Count + 1, collatz_steps(M, Steps, NewCount).
collatz_steps(N, Steps, Count) :- M is 3 * N + 1, NewCount is Count + 1, collatz_steps(M, Steps, NewCount).
