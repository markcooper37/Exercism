//
// This is only a SKELETON file for the 'Matrix' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class Matrix {
  constructor(matrix) {
    this._rows = matrix.split('\n').map(row => row.split(' ').map(digit => Number(digit)));
    if (this._rows.length === 0) {
      this._columns = [];
    } else {
      let columns = [];
      for (let i = 0; i < this._rows[0].length; i++) {
        let column = [];
        for (let j = 0; j < this._rows.length; j++) {
          column.push(this._rows[j][i]);
        };
        columns.push(column);
      };
      this._columns = columns;
    };
  }

  get rows() {
    return this._rows;
  }

  get columns() {
    return this._columns;
  }
}
