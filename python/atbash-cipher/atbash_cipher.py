def encode(plain_text):
    characters = [character for character in plain_text.replace(" ", "").lower() if character.isalnum()]
    return " ".join("".join([chr(219-ord(ch)) if ch.isalpha() else ch for ch in characters][i:i + 5]) for i in range(0, len(characters), 5))


def decode(ciphered_text):
    return "".join([chr(219-ord(ch)) if ch.isalpha() else ch for ch in ciphered_text.replace(" ", "")])
