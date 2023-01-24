export const saddlePoints = (matrix) => {
  let saddlePoints = [];
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[0].length; j++) {
      if (smallestInColumn(matrix, i, j) && largestInRow(matrix, i, j)) {
        saddlePoints.push({row: i + 1, column: j + 1})
      };
    };
  };
  return saddlePoints;
};

const smallestInColumn = (matrix, row, column) => {
  return matrix.map(matrixRow => matrixRow[column]).every(value => value >= matrix[row][column]);
}

const largestInRow = (matrix, row, column) => {
  return matrix[row].every(value => value <= matrix[row][column]);
}
