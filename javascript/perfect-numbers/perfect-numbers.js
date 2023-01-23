export const classify = (number) => {
  if (number < 1) {
    throw new Error('Classification is only possible for natural numbers.');
  };
  let factorSum = 0;
  for (let i = 1; i <= number / 2; i++) {
    if (number % i === 0) {
      factorSum += i;
    };
  };
  if (factorSum < number) {
    return 'deficient';
  } else if (factorSum === number) {
    return 'perfect';
  } else {
    return 'abundant';
  };
};
