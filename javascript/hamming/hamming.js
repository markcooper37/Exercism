//
// This is only a SKELETON file for the 'Hamming' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const compute = (strand1, strand2) => {
  if (strand1.length !== strand2.length) {
    throw new Error('strands must be of equal length');
  };
  return strand1.split('').filter((letter, index) => strand2[index] !== letter).length;
};
