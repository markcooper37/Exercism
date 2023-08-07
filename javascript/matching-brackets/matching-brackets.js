export const isPaired = (brackets) => {
  let bracketMap = new Map([['[', ']'], ['{', '}'], ['(', ')']]);
  let openBrackets = [];
  for (let char of brackets) {
    if (['[', '{', '('].includes(char)) {
      openBrackets.push(char);
    } else if ([']', '}', ')'].includes(char)) {
      if (openBrackets.length === 0) {
        return false;
      }
      let openBracket = openBrackets.pop();
      if (bracketMap.get(openBracket) !== char) {
        return false;
      }
    }
  }
  return openBrackets.length === 0;
};
