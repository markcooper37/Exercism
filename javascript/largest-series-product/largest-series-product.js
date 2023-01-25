export const largestProduct = (digits, span) => {
  if (span > digits.length) {
    throw new Error('Span must be smaller than string length');
  } else if (digits.match(/[^0-9]/)) {
    throw new Error('Digits input must only contain digits');
  } else if (span < 0) {
    throw new Error('Span must be greater than zero');
  }
  let [maxProduct, numbers] = [0, [...digits].map(digit => Number(digit))];
  for (let i = 0; i <= digits.length - span; i++) {
    let product = numbers.slice(i, i + span).reduce((prod, value) => prod * value, 1);
    maxProduct = product > maxProduct ? product : maxProduct;
  };
  return maxProduct;
};
