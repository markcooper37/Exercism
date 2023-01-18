export class Clock {
  constructor(hours = 0, minutes = 0) {
    this.minutes = 60 * hours + minutes;
  }

  toString() {
    let minutes = ((this.minutes % 60) + 60) % 60
    let hours = ((Math.floor(this.minutes / 60) % 24) + 24) % 24;
    return `${('0' + hours).slice(-2)}:${('0' + minutes).slice(-2)}`;
  }

  plus(minutes) {
    return new Clock(0, this.minutes + minutes);
  }

  minus(minutes) {
    return new Clock(0, this.minutes - minutes);
  }

  equals(clock2) {
    return this.toString() === clock2.toString();
  }
}
