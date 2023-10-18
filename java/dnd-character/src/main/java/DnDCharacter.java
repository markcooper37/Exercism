import java.util.List;
import java.util.ArrayList;

class DnDCharacter {

    private int strength;
    private int dexterity;
    private int constitution;
    private int intelligence;
    private int wisdom;
    private int charisma;
    private int hitpoints;

    public DnDCharacter() {
        this.strength = ability(rollDice());
        this.dexterity = ability(rollDice());
        this.constitution = ability(rollDice());
        this.intelligence = ability(rollDice());
        this.wisdom = ability(rollDice());
        this.charisma = ability(rollDice());
        this.hitpoints = modifier(this.constitution) + 10;
    }

    int ability(List<Integer> scores) {
        int min = scores.get(0);
        int total = 0;
        for (int i = 1; i <= 3; i++) {
            int value = scores.get(i);
            if (value < min) {
                total += min;
                min = value;
            } else {
                total += value;
            }
        }
        return total;
    }

    List<Integer> rollDice() {
        List<Integer> dice = new ArrayList<>();
        for (int i = 0; i < 4; i++) {
            dice.add((int) Math.floor(Math.random() * 6 + 1));
        }
        return dice;
    }

    int modifier(int input) {
        return input / 2 - 5;
    }

    int getStrength() {
        return this.strength;
    }

    int getDexterity() {
        return this.dexterity;
    }

    int getConstitution() {
        return this.constitution;
    }

    int getIntelligence() {
        return this.intelligence;
    }

    int getWisdom() {
        return this.wisdom;
    }

    int getCharisma() {
        return this.charisma;
    }

    int getHitpoints() {
        return this.hitpoints;
    }
}
