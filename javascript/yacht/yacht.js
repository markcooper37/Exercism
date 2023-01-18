//
// This is only a SKELETON file for the 'Yacht' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const score = (dice, category) => {
  dice.sort((die1, die2) => die1 - die2);
  switch (category) {
    case 'ones':
      return dice.filter(die => die === 1).length;
    case 'twos':
      return dice.filter(die => die === 2).length * 2;
    case 'threes':
      return dice.filter(die => die === 3).length * 3;
    case 'fours':
      return dice.filter(die => die === 4).length * 4;
    case 'fives':
      return dice.filter(die => die === 5).length * 5;
    case 'sixes':
      return dice.filter(die => die === 6).length * 6;
    case 'full house':
      let [lowerLength, higherLength] = [dice.filter(die => die === dice[0]).length, dice.filter(die => die === dice[4]).length];
      return (lowerLength + higherLength === 5 && [2, 3].includes(lowerLength)) * (dice[0] * lowerLength + dice[4] * higherLength);
    case 'four of a kind':
      return (dice.filter(die => die === dice[2]).length >= 4) * dice[2] * 4;
    case 'little straight':
      return (dice[0] === 1 && isStraight(dice)) * 30;
    case 'big straight':
      return (dice[0] === 2 && isStraight(dice)) * 30;
    case 'choice':
      return dice.reduce((accumulator, die) => accumulator + die, 0);
    case 'yacht':
      return dice.every(die => die === dice[0]) * 50;
  };
};

function isStraight(dice) {
  for (let i = 0; i < dice.length - 1; i++) {
    if (dice[i+1] - dice[i] != 1) {
      return false;
    };
  };
  return true;
}
