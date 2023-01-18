//
// This is only a SKELETON file for the 'All Your Base' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const convert = (digits, inputBase, outputBase) => {
  if (inputBase <= 1) {
    throw new Error('Wrong input base');
  } else if (outputBase <= 1) {
    throw new Error('Wrong output base');
  };
  if (digits.length === 0) {
    throw new Error('Input has wrong format');
  } else if (digits[0] === 0) {
    if (digits.length === 1) {
      return [0];
    } else {
      throw new Error('Input has wrong format');
    };
  };
  let number = 0;
  for (let i = 0; i < digits.length; i++) {
    if (digits[digits.length-i-1] < 0 || digits[digits.length-i-1] >= inputBase) {
      throw new Error('Input has wrong format');
    };
    number += digits[digits.length-i-1] * (inputBase ** i);
  };
  let output = [];
  while (number > 0) {
    output.unshift(number % outputBase);
    number = (number - (number % outputBase)) / outputBase;
  };
  return output
};
