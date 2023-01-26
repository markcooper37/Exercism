def classify(number):
    """ A perfect number equals the sum of its positive divisors.

    :param number: int a positive integer
    :return: str the classification of the input integer
    """

    if number <= 0:
        raise ValueError('Classification is only possible for positive integers.')
    sum = 0
    for value in range(1, number // 2 + 1):
        if number % value == 0:
            sum += value
    return 'perfect' if sum == number else 'abundant' if sum > number else 'deficient'
