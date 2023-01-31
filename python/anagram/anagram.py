def find_anagrams(word, candidates):
    word_lower = word.lower()
    ordered_word = sorted(list(word_lower))
    anagrams = []
    for candidate in candidates:
        candidate_lower = candidate.lower()
        if sorted(list(candidate_lower)) == ordered_word and candidate_lower != word_lower:
            anagrams.append(candidate)
    return anagrams
        
