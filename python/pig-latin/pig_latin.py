def translate(text):
    words = []
    for word in text.split():
        if len(word) >= 2 and word[0:2] in ["xr", "yt"]:
            words.append(word + "ay")
            continue
        for pos in range(len(word)):
            if word[pos] == "q" and pos < len(word) - 1 and word[pos + 1] == "u":
                pos += 2
            elif word[pos] not in ["a", "e", "i", "o", "u"] and (pos == 0 or word[pos] != "y"):
                continue
            break
        words.append(word[pos:] + word[:pos] + "ay")
    return " ".join(words)
