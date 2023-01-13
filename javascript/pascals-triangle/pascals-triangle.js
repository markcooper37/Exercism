//
// This is only a SKELETON file for the 'Pascals Triangle' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const rows = (numberOfRows) => {
  let rows = [];
  for (let i = 0; i < numberOfRows; i++) {
    if (i === 0) {
      rows.push([1]);
      continue;
    }
    let [row, rowAbove] = [[], [0, ...rows[i-1], 0]];
    for (let j = 0; j <= i; j++) {
      row.push(rowAbove[j] + rowAbove[j + 1]);
    };
    rows.push(row);
  };
  return rows;
};
