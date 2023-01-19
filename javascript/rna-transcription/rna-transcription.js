//
// This is only a SKELETON file for the 'RNA Transcription' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const toRna = (strand) => {
  let nucleotideMap = {G: 'C', C: 'G', T: 'A', A: 'U'};
  return strand.split('').map(value => nucleotideMap[value]).join('');
};
