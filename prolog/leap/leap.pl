leap(Year) :- 0 =:= mod(Year, 400).

leap(Year) :- 0 =:= mod(Year, 4), 0 =\= mod(Year, 100).
