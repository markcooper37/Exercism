import math


def largest_product(series, size):
    if size > len(series):
        raise ValueError("span must be smaller than string length")
    if size < 0:
        raise ValueError("span must not be negative")
    if len(series) > 0 and not series.isnumeric():
        raise ValueError("digits input must only contain digits")
    maximum = 0
    for i in range(len(series) - size + 1):
        maximum = max(math.prod([int(x) for x in list(series[i:i + size])]), maximum)
    return maximum
