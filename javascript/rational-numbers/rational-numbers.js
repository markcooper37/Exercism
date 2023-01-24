const gcd = (int1, int2) => !int2 ? int1 : gcd(int2, int1 % int2);

export class Rational {
  constructor(num, denom) {
    [this.num, this.denom] = denom >= 0 ? [num, denom] : [-num, -denom];
  }

  add(rational2) {
    return new Rational(this.num * rational2.denom + this.denom * rational2.num, this.denom * rational2.denom).reduce();
  }

  sub(rational2) {
    return new Rational(this.num * rational2.denom - this.denom * rational2.num, this.denom * rational2.denom).reduce();
  }

  mul(rational2) {
    return new Rational(this.num * rational2.num, this.denom * rational2.denom).reduce();
  }

  div(rational2) {
    return new Rational(this.num * rational2.denom, this.denom * rational2.num).reduce();
  }

  abs() {
    return new Rational(Math.abs(this.num), Math.abs(this.denom)).reduce();
  }

  exprational(exponent) {
    return exponent >= 0 ? new Rational(this.num ** exponent, this.denom ** exponent).reduce() : new Rational(this.denom ** -exponent, this.num ** -exponent).reduce();
  }

  expreal(number) {
    return (number ** (1/this.denom)) ** this.num;
  }

  reduce() {
    return new Rational(this.num/gcd(this.num, this.denom), this.denom/gcd(this.num, this.denom));
  }
}
