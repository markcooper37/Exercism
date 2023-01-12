/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

export function cookingStatus(remainingTime) {
  switch (remainingTime) {
    case undefined:
      return 'You forgot to set the timer.';
    case 0:
      return 'Lasagna is done.';
    default:
      return 'Not done, please wait.';
  };
}

export function preparationTime(layers, timePerLayer = 2) {
  return layers.length * timePerLayer;
}

export function quantities(layers) {
  let amounts = {
    noodles: 0,
    sauce: 0,
  };
  for (let i = 0; i < layers.length; i++) {
    if (layers[i] === 'sauce') {
      amounts.sauce += 0.2;
    } else if (layers[i] === 'noodles') {
      amounts.noodles += 50;
    };
  };
  return amounts;
}

export function addSecretIngredient(ingredients1, ingredients2) {
  ingredients2.push(ingredients1[ingredients1.length - 1]);
}

export function scaleRecipe(recipe, portions) {
  let newRecipe = {};
  for (let key in recipe) {
    newRecipe[key] = recipe[key] / 2 * portions;
  };
  return newRecipe;
}
