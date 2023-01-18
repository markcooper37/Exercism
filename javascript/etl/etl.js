export const transform = (scores) => {
  let newScores = {};
  for (let score in scores) {
    for (const letter of scores[score]) {
      newScores[letter.toLowerCase()] = Number(score);
    };
  };
  return newScores;
};
