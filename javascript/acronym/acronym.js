export const parse = (phrase) => {
  return phrase.replaceAll('-', ' ').split(/[ ]+/).map(word => word.match(/[a-zA-Z]/)[0].toUpperCase()).join('');
};
