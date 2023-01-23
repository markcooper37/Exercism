export class Series {
  constructor(series) {
    if (series.length === 0) {
      throw new Error('series cannot be empty');
    }
    this._series = series.split('').map(value => Number(value));
  }

  slices(sliceLength) {
    if (sliceLength === 0) {
      throw new Error('slice length cannot be zero');
    } else if (sliceLength < 0) {
      throw new Error('slice length cannot be negative');
    } else if (sliceLength > this._series.length) {
      throw new Error('slice length cannot be greater than series length');
    };
    let slices = [];
    for (let i = 0; i <= this._series.length - sliceLength; i++) {
      slices.push(this._series.slice(i, i + sliceLength));
    };
    return slices;
  }
}
