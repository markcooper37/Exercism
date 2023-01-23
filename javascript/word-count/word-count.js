export const countWords = (phrase) => {
  return phrase.match(/[a-z]+'[a-z]+|[a-z]+|[0-9]+/gi).reduce((wordCounts, word) => {
    let lowerCaseWord = word.toLowerCase();
    wordCounts[lowerCaseWord] = (!wordCounts[lowerCaseWord] ? 1 : wordCounts[lowerCaseWord] + 1)
    return wordCounts;
  }, {});
};
