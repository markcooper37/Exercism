hey(Sentence, Response) :- is_empty(Sentence), Response = "Fine. Be that way!".
hey(Sentence, Response) :- is_question(Sentence), is_shouting(Sentence), Response = "Calm down, I know what I'm doing!".
hey(Sentence, Response) :- is_question(Sentence), Response = "Sure.".
hey(Sentence, Response) :- is_shouting(Sentence), Response = "Whoa, chill out!".
hey(_, Response) :- Response = "Whatever.".

is_empty(Sentence) :- string_chars(Sentence, CharsList), include(is_not_space, CharsList, Chars), length(Chars, Length), Length = 0.

is_question(Sentence) :- string_chars(Sentence, CharsList), include(is_not_space, CharsList, Chars), last(Chars, Last), Last = '?'.

is_shouting(Sentence) :- string_chars(Sentence, CharsList), include(is_alpha, CharsList, Chars), string_chars(Letters, Chars), string_upper(Letters, Lower), Letters = Lower.

is_not_space(Char) :- not(is_space(Char)).
