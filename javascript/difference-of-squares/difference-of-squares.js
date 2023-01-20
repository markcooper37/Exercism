export class Squares {
  constructor(value) {
    this._value = value;
  }

  get sumOfSquares() {
    return this._value * (this._value + 1) * (2 * this._value + 1) / 6;
  }

  get squareOfSum() {
    return ((this._value * (this._value + 1)) / 2) ** 2;
  }

  get difference() {
    return this.squareOfSum - this.sumOfSquares;
  }
}
