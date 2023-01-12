//
// This is only a SKELETON file for the 'Prime Factors' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const primeFactors = (number) => {
  let factors = [];
  let factor = 2;
  while (number > 1) {
    if (number % factor === 0) {
      factors.push(factor);
      number /= factor;
    } else {
      factor++;
    };
  };
  return factors;
};
