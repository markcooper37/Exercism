class Allergies {

    def allergies = []

    private static final allergens = ["eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"]

    Allergies(int score) {
        for (def i = 0; i <= 7; i++) {
            if (score & 1 << i) {
                allergies.add(allergens[i])
            }
        }
    }

    def allergicTo(String substance) {
        return allergies.contains(substance)
    }

    def list() {
        return allergies
    }

}