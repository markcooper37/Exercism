//
// This is only a SKELETON file for the 'Sum Of Multiples' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const sum = (numberList, limit) => {
  const multiples = new Map();
  for (const number of numberList) {
    if (number === 0) {
      continue;
    };
    for (let i = number; i < limit; i += number) {
      multiples.set(i, true);
    };
  };
  let sum = 0;
  for (const [key,] of multiples) {
    sum += key;
  };
  return sum;
};
