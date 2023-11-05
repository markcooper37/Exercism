isogram(Phrase) :- alphabetic_chars(Phrase, Chars), is_set(Chars).

alphabetic_chars(Phrase, Chars) :- string_lower(Phrase, PhraseLower), string_chars(PhraseLower, CharsList), include(is_alpha, CharsList, Chars).
