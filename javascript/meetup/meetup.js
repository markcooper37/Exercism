const days = new Map([['Sunday', 0], ['Monday', 1], ['Tuesday', 2], ['Wednesday', 3], ['Thursday', 4], ['Friday', 5], ['Saturday', 6]]);

export const meetup = (year, month, descriptor, day) => {
  let start = new Date(year, month - 1, 1);
  let offset;
  switch (descriptor) {
    case 'last':
      let end = new Date(year, month, 0);
      let monthLength = end.getDate();
      offset = monthLength - 1 - ((end.getDay() - days.get(day) + 7) % 7);
      break;
    case 'teenth':
      let teen = new Date(year, month - 1, 13);
      offset = 12 + ((days.get(day) - teen.getDay() + 7) % 7);
      break;
    case 'first':
      offset = (days.get(day) - start.getDay() + 7) % 7;
      break;
    case 'second':
      offset = 7 + ((days.get(day) - start.getDay() + 7) % 7);
      break;
    case 'third':
      offset = 14 + ((days.get(day) - start.getDay() + 7) % 7);
      break;
    case 'fourth':
      offset = 21 + ((days.get(day) - start.getDay() + 7) % 7);
  }
  return new Date(year, month - 1, 1 + offset);
};
