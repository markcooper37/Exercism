export function triplets({ sum, minFactor = 1, maxFactor = sum }) {
  let triplets = [];
  for (let a = minFactor; a <= maxFactor - 2; a++) {
    for (let b = a + 1; b <= maxFactor - 1; b++) {
      for (let c = b + 1; c <= maxFactor; c++) {
        if (a ** 2 + b ** 2 === c ** 2 && a + b + c === sum) {
          triplets.push(new Triplet(a, b, c));
        };
      };
    };
  };
  return triplets;
}

class Triplet {
  constructor(a, b, c) {
    [this._a, this._b, this._c] = [a, b, c];
  }

  toArray() {
    return [this._a, this._b, this._c];
  }
}
