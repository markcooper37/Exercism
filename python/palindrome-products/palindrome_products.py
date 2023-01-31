def largest(max_factor, min_factor = 0):
    """Given a range of numbers, find the largest palindromes which
       are products of two numbers within that range.

    :param min_factor: int with a default value of 0
    :param max_factor: int
    :return: tuple of (palindrome, iterable).
             Iterable should contain both factors of the palindrome in an arbitrary order.
    """

    if min_factor > max_factor:
        raise ValueError("min must be <= max")
    value, factors = -1, []
    for i in range(max_factor, min_factor - 1, -1):
        if i ** 2 < value:
            break
        for j in range(i, min_factor - 1, -1):
            digits = [digit for digit in str(i * j)]
            if digits == digits[::-1]:
                if value == i * j:
                    factors.append([i, j])
                elif value < i * j:
                    value = i * j
                    factors = [[i, j]]
    return ((value if value != -1 else None), factors)
            


def smallest(max_factor, min_factor = 0):
    """Given a range of numbers, find the smallest palindromes which
    are products of two numbers within that range.

    :param min_factor: int with a default value of 0
    :param max_factor: int
    :return: tuple of (palindrome, iterable).
    Iterable should contain both factors of the palindrome in an arbitrary order.
    """

    if min_factor > max_factor:
        raise ValueError("min must be <= max")
    value, factors = max_factor ** 2 + 1, []
    for i in range(min_factor, max_factor + 1):
        if i ** 2 > value:
            break
        for j in range(i, max_factor + 1):
            digits = [digit for digit in str(i * j)]
            if digits == digits[::-1]:
                if value == i * j:
                    factors.append([i, j])
                elif value > i * j:
                    value = i * j
                    factors = [[i, j]]
    return ((value if value != max_factor ** 2 + 1 else None), factors)
