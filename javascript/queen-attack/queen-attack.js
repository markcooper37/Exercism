//
// This is only a SKELETON file for the 'Queen Attack' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class QueenAttack {
  constructor({
    black: [blackRow, blackColumn] = [],
    white: [whiteRow, whiteColumn] = [],
  } = {}) {
    if (blackRow < 0 || blackRow > 7 || blackColumn < 0 || blackColumn > 7 || whiteRow < 0 || whiteRow > 7 || whiteColumn < 0 || whiteColumn > 7) {
      throw new Error('Queen must be placed on the board');
    }
    if (blackRow == undefined) {
      this.black = [0, 3];
    } else {
      this.black = [blackRow, blackColumn];
    }
    if (whiteRow == undefined) {
      this.white = [7, 3];
    } else {
      this.white = [whiteRow, whiteColumn];
    }
    if (this.black[0] == this.white[0] && this.black[1] == this.white[1]) {
      throw new Error('Queens cannot share the same space');
    }
  }

  toString() {
    let board = '';
    for (let i = 0; i < 8; i++) {
       let row = '';
      for (let j = 0; j < 8; j++) {
        if (i == this.black[0] && j == this.black[1]) {
          row += 'B';
        } else if (i == this.white[0] && j == this.white[1]) {
          row += 'W';
        } else {
          row += '_';
        }
        if (j < 7) {
          row += ' ';
        }
      }
      board += row;
      if (i < 7) {
        board += '\n';
      }
    }
    return board;
  }

  get canAttack() {
    return this.black[0] == this.white[0] || this.black[1] == this.white[1] || this.black[0] + this.black[1] == this.white[0] + this.white[1] || this.black[0] - this.black[1] == this.white[0] - this.white[1];
  }
}
