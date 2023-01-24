export const isValid = (number) => {
  let formattedNumber = [...number].filter((value, index) => {
    return value.match(/[0-9]/) || value === 'X' && index === number.length - 1;
  }).map(value => value === 'X' ? 10 : value);
  return formattedNumber.length !== 10 ? false : formattedNumber.reduce((sum, value, index) => sum + value * (10 - index), 0) % 11 === 0;
};
