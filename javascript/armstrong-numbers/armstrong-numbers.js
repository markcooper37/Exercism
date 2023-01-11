//
// This is only a SKELETON file for the 'Armstrong Numbers' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const isArmstrongNumber = (value) => {
  let sum = 0;
  let length = value.toString().length;
  for (let i = 0; i < length; i++) {
    sum += (((value - (value % 10 ** i)) / (10 ** i)) % 10) ** length;
  }
  return sum === value;
};
