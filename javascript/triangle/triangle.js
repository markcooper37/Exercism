//
// This is only a SKELETON file for the 'Triangle' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class Triangle {
  constructor(...sides) {
    this._sides = sides.sort((side1, side2) => side1 - side2);
    this._isTriangle = this._sides[0] > 0 && this._sides[0] + this._sides[1] > this._sides[2];
  }

  get isEquilateral() {
    return this._isTriangle && (this._sides[0] === this._sides[1] && this._sides[1] === this._sides[2]);
  }

  get isIsosceles() {
    return this._isTriangle && (this._sides[0] === this._sides[1] || this._sides[1] === this._sides[2]);
  }

  get isScalene() {
    return this._isTriangle && (this._sides[0] !== this._sides[1] && this._sides[1] !== this._sides[2]);
  }
}
