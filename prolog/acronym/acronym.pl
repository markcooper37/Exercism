abbreviate(Sentence, Acronym) :- get_words(Sentence, Words), maplist(get_initial, Words, Initials), atomics_to_string(Initials, Acronym).

get_words(Sentence, Words) :- split_string(Sentence, "- ", "_ ", SubStrings), include(is_string_not_empty, SubStrings, Words).

is_string_not_empty(Str) :- not(string_length(Str, 0)).

get_initial(Word, Initial) :- string_upper(Word, Upper), sub_string(Upper, 0, 1, _, Initial).
