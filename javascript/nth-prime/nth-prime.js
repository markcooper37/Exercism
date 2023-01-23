//
// This is only a SKELETON file for the 'Nth Prime' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const prime = (n) => {
  if (n === 0) {
    throw new Error('there is no zeroth prime');
  };
  let [number, primes] = [2, []];
  while (primes.length < n) {
    let isPrime = true;
    for (const prime of primes) {
      if (number % prime === 0) {
        isPrime = false;
        break;
      };
    };
    if (isPrime) {
      primes.push(number);
    };
    number++;
  };
  return primes[primes.length - 1];
};
