export const knapsack = (maximumWeight, items) => {
  let maximumValue = 0;
  for (let i = 0; i < 2 ** items.length; i++) {
    let [weight, value] = [0, 0];
    for (let j = 0; j < items.length; j++) {
      if (i >> j & 1) {
        weight += items[j].weight;
        value += items[j].value;
      };
    };
    if (weight <= maximumWeight && value > maximumValue) {
      maximumValue = value;
    };
  };
  return maximumValue;
};
