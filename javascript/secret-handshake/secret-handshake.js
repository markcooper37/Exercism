export const commands = (number) => {
  let sequence = ['wink', 'double blink', 'close your eyes', 'jump'].filter((_, i) => number >> i & 1);
  return (number >> 4 & 1 ? sequence.reverse() : sequence);
};
