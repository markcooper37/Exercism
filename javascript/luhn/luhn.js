export const valid = (number) => {
  let trimmedNumber = number.replace(/ /g, '');
  if (!/^\d+$/.test(trimmedNumber) || trimmedNumber.length <= 1) {
    return false;
  };
  return trimmedNumber.split('').reduce((accumulator, value, i) => {
    return accumulator + ((i + trimmedNumber.length) % 2 !== 0 ? Number(value) : Number(value) * 2 >= 10 ? Number(value) * 2 - 9 : Number(value) * 2);
  }, 0) % 10 === 0;
};
