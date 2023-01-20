export class Palindromes {
  static generate(factors) {
    if (factors.minFactor > factors.maxFactor) {
      throw new Error('min must be <= max');
    };
    let palindromesObject = {};
    for (let i = factors.minFactor; i <= factors.maxFactor; i++) {
      for (let j = i; j <= factors.maxFactor; j++) {
        if (isPalindrome(i * j)) {
          if (!palindromesObject[i * j]) {
            palindromesObject[i * j] = [[i, j]];
          } else {
            palindromesObject[i * j] = palindromesObject[i * j].concat([[i, j]]);
          };
        };
      };
    };
    let palindromesArray = [];
    for (let palindrome in palindromesObject) {
      palindromesArray.push({value: Number(palindrome), factors: palindromesObject[palindrome]});
    };
    palindromesArray.sort((p1, p2) => {
      return p1.value - p2.value;
    });
    let extremes = {smallest: {value: null, factors: []}, largest: {value: null, factors: []}};
    if (palindromesArray.length > 0) {
      extremes.smallest = palindromesArray[0];
      extremes.largest = palindromesArray[palindromesArray.length - 1];
    };
    return extremes;
  }
}

function isPalindrome(number) {
  let stringNumber = number.toString()
  for (let i = 0; i < stringNumber.length / 2; i++) {
    if (stringNumber[i] !== stringNumber[stringNumber.length - i - 1]) {
      return false;
    };
  };
  return true;
}
