def prime(number):
    if number == 0:
        raise ValueError("there is no zeroth prime")
    primes, test = [2], 3
    while len(primes) < number:
        is_prime = True
        for prime in primes:
            if test % prime == 0:
                is_prime = False
                break
        if is_prime:
            primes.append(test)
        test += 1
    return primes[-1]
