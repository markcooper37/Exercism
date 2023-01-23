export const isIsogram = (phrase) => {
  let trimmedPhrase = phrase.toLowerCase().split('').filter((value) => value.match(/[a-z]/));
  return new Set(trimmedPhrase).size === trimmedPhrase.length;
};
