def is_valid(isbn):
    digits = [digit for digit in isbn if digit.isalnum()]
    if len(digits) != 10 or [digit for index, digit in enumerate(digits) if digit.isnumeric() or index == len(digits) - 1 and digit == 'X'] != digits:
        return False
    return sum([(int(value) if value != 'X' else 10) * (index + 1) for index, value in enumerate(digits)]) % 11 == 0
