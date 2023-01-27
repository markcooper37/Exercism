def rotate(text, key):
    return ''.join([value if not value.isalpha() else chr((ord(value) - 65 - 32 * value.islower() + key) % 26 + 65 + 32 * value.islower()) for value in text])
