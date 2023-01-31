def encode(numbers):
    new_numbers = []
    for number in numbers:
        new_number = [number % 128]
        number >>= 7
        while number > 0:
            new_number.insert(0, number % 128 | (1 << 7))
            number >>= 7
        new_numbers.extend(new_number)
    return new_numbers


def decode(bytes_):
    if bytes_[-1] >> 7 & 1:
        raise ValueError("incomplete sequence")
    numbers, number = [], 0
    for value in bytes_:
        number <<= 7
        number += value % 128
        if not value & (1 << 7):
            numbers.append(number)
            number = 0
    return numbers
