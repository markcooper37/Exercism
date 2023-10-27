convert(N, Sounds) :- N mod 105 =:= 0 -> Sounds = "PlingPlangPlong".
convert(N, Sounds) :- N mod 15 =:= 0 -> Sounds = "PlingPlang".
convert(N, Sounds) :- N mod 21 =:= 0 -> Sounds = "PlingPlong".
convert(N, Sounds) :- N mod 35 =:= 0 -> Sounds = "PlangPlong".
convert(N, Sounds) :- N mod 3 =:= 0 -> Sounds = "Pling".
convert(N, Sounds) :- N mod 5 =:= 0 -> Sounds = "Plang".
convert(N, Sounds) :- N mod 7 =:= 0 -> Sounds = "Plong".
convert(N, Sounds) :- number_string(N, Sounds).
