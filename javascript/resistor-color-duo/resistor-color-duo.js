//
// This is only a SKELETON file for the 'Resistor Color Duo' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const decodedValue = ([name1, name2,]) => {
  return 10 * resistorValue(name1) + resistorValue(name2);
};

export const COLORS = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white'];

export function resistorValue(name) {
  return COLORS.indexOf(name);
}
