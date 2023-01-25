export const primes = (limit) => {
  let [output, composites] = [[], new Set()];
  for (let i = 2; i <= limit; i++) {
    if (!composites.has(i)) {
      output.push(i);
      for (let j = 2; j * i <= limit; j++) {
        composites.add(j * i);
      };
    };
  };
  return output;
};
