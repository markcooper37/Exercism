def abbreviate(words):
    return "".join([next(char for char in word if char.isalpha()).upper() for word in words.replace("-", " ").split()])
