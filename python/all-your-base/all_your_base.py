def rebase(input_base, digits, output_base):
    if input_base < 2:
        raise ValueError('input base must be >= 2')
    if output_base < 2:
        raise ValueError('output base must be >= 2')
    number = 0
    for index, value in enumerate(digits[::-1]):
        if value < 0 or value >= input_base:
            raise ValueError('all digits must satisfy 0 <= d < input base')
        number += value * input_base ** index
    if number == 0:
        return [0]
    new_digits = []
    while number > 0:
        new_digits.insert(0, (number % output_base))
        number //= output_base
    return new_digits
