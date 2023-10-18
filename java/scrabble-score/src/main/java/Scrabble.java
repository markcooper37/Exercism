class Scrabble {

    private final int score;
    private static int[] letterValues = {1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10};

    Scrabble(String word) {
        this.score = word
                .toLowerCase()
                .chars()
                .map(c -> c - 'a')
                .map(i -> letterValues[i])
                .sum();
    }

    int getScore() {
        return this.score;
    }

}
