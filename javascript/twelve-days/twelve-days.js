const days = ['first', 'second', 'third', 'fourth', 'fifth', 'sixth', 'seventh', 'eighth', 'ninth', 'tenth', 'eleventh', 'twelfth'];

const items = ['twelve Drummers Drumming,', 'eleven Pipers Piping,', 'ten Lords-a-Leaping,', 'nine Ladies Dancing,', 'eight Maids-a-Milking,', 'seven Swans-a-Swimming,', 'six Geese-a-Laying,', 'five Gold Rings,', 'four Calling Birds,', 'three French Hens,', 'two Turtle Doves, and', 'a Partridge in a Pear Tree.'];

export const recite = (start, end = start) => {
  let verses = [];
  for (let number = start; number <= end; number++) {
    verses.push(verse(number));
  }
  return verses.join('\n');
};

function verse(number) {
  return 'On the ' + days[number - 1] + ' day of Christmas my true love gave to me: ' + items.slice(12 - number).join(' ') + '\n';
}