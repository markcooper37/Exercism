import re

def count_words(sentence):
    matches = re.findall(r'\b\w+[\'-]?\w*\b', sentence.lower().replace("_", " "))
    return {word: matches.count(word) for word in matches}
