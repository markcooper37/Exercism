export class Allergies {
  static ALLERGENS = ['eggs', 'peanuts', 'shellfish', 'strawberries', 'tomatoes', 'chocolate', 'pollen', 'cats'];
  
  constructor(score) {
    this._allergies = [];
    for (let i = 0; i < 8; i++) {
      if ((score % (2 ** (i + 1))) / (2 ** i) >= 1) {
        this._allergies.push(Allergies.ALLERGENS[i]);
      };
    };
  }

  list() {
    return this._allergies;
  }

  allergicTo(allergen) {
    return this._allergies.includes(allergen);
  }
}
