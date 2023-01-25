export const flatten = (input) => {
  let output = [];
  for (const item of input) {
    if (Array.isArray(item)) {
      output = output.concat(flatten(item));
    } else if (item !== null) {
      output.push(item);
    };
  };
  return output;
};
