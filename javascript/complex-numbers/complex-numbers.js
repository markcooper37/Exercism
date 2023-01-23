export class ComplexNumber {
  constructor(real, imag) {
    [this._real, this._imag] = [real, imag];
  }

  get real() {
    return this._real;
  }

  get imag() {
    return this._imag;
  }

  add(num2) {
    return new ComplexNumber(this.real + num2.real, this.imag + num2.imag);
  }

  sub(num2) {
    return new ComplexNumber(this.real - num2.real, this.imag - num2.imag);
  }

  div(num2) {
    let realPart = (this.real * num2.real + this.imag * num2.imag) / (num2.real ** 2 + num2.imag ** 2);
    let imagPart = (this.imag * num2.real - this.real * num2.imag) / (num2.real ** 2 + num2.imag ** 2);
    return new ComplexNumber(realPart, imagPart);
  }

  mul(num2) {
    return new ComplexNumber(this.real * num2.real - this.imag * num2.imag, this.real * num2.imag + this.imag * num2.real);
  }

  get abs() {
    return Math.sqrt(this.real ** 2 + this.imag ** 2);
  }

  get conj() {
    return new ComplexNumber(this.real, this.imag === 0 ? 0 : -this.imag);
  }

  get exp() {
    return new ComplexNumber(Math.E ** this.real * Math.cos(this.imag), Math.E ** this.real * Math.sin(this.imag));
  }
}
