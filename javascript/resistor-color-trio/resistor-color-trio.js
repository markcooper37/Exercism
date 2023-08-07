const colorValues = new Map([
  ['black', 0], 
  ['brown', 1], 
  ['red', 2], 
  ['orange', 3], 
  ['yellow', 4], 
  ['green', 5], 
  ['blue', 6], 
  ['violet', 7], 
  ['grey', 8], 
  ['white', 9]
]);

export class ResistorColorTrio {
  constructor(colors) {
    let value = (10 * colorValues.get(colors[0]) + colorValues.get(colors[1])) * (10 ** colorValues.get(colors[2]));
    if (isNaN(value)) throw new Error('invalid color');
    this.value = value;
  }

  get label() {
    if (this.value >= 1000) {
      return `Resistor value: ${this.value/1000} kiloohms`;
    } else {
      return `Resistor value: ${this.value} ohms`;
    }
  }
}
