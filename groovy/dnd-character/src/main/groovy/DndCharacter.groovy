class DndCharacter {

    def final strength
    def final dexterity
    def final constitution
    def final intelligence
    def final wisdom
    def final charisma
    def final hitpoints

    DndCharacter() {
        this.strength = ability();
        this.dexterity = ability();
        this.constitution = ability();
        this.intelligence = ability();
        this.wisdom = ability();
        this.charisma = ability();
        this.hitpoints = modifier(this.constitution) + 10;
    }

    def ability() {
        List<Integer> dice = new ArrayList<>();
        for (int i = 0; i < 4; i++) {
            dice.add((int) Math.floor(Math.random() * 6 + 1));
        }
        int min = dice.get(0);
        int total = 0;
        for (int i = 1; i <= 3; i++) {
            int value = dice.get(i);
            if (value < min) {
                total += min;
                min = value;
            } else {
                total += value;
            }
        }
        return total;
    }

    def modifier(int value) {
        return value.intdiv(2) - 5
    }

}
