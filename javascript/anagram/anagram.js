export const findAnagrams = (word, candidates) => {
  return candidates.filter(candidate => [...word.toLowerCase()].sort().join('') === [...candidate.toLowerCase()].sort().join('') && word.toLowerCase() !== candidate.toLowerCase());
};
