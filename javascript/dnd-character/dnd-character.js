//
// This is only a SKELETON file for the 'D&D Character' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const abilityModifier = (score) => {
  if (score < 3) {
    throw new Error('Ability scores must be at least 3');
  } else if (score > 18) {
    throw new Error('Ability scores can be at most 18');
  };
  return Math.floor((score - 10) / 2);
};

export class Character {
  static rollAbility() {
    let throws = [];
    for (let i = 0; i < 4; i++) {
      throws.push(Math.floor(Math.random() * 6 + 1));
    }
    throws.sort();
    return throws[1] + throws[2] + throws[3];
  }

  get strength() {
    if (this.strengthScore === undefined) {
      this.strengthScore = Character.rollAbility()
    }
    return this.strengthScore;
  }

  get dexterity() {
    if (this.dexterityScore === undefined) {
      this.dexterityScore = Character.rollAbility()
    }
    return this.dexterityScore;
  }

  get constitution() {
    if (this.constitutionScore === undefined) {
      this.constitutionScore = Character.rollAbility()
    }
    return this.constitutionScore;
  }

  get intelligence() {
    if (this.intelligenceScore === undefined) {
      this.intelligenceScore = Character.rollAbility()
    }
    return this.intelligenceScore;
  }

  get wisdom() {
    if (this.wisdomScore === undefined) {
      this.wisdomScore = Character.rollAbility()
    }
    return this.wisdomScore;
  }

  get charisma() {
    if (this.charismaScore === undefined) {
      this.charismaScore = Character.rollAbility()
    }
    return this.charismaScore;
  }

  get hitpoints() {
    return 10 + abilityModifier(this.constitutionScore);
  }
}
