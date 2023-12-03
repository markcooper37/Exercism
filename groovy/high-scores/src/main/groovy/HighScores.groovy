class HighScores {

    static List<Integer> scores

    HighScores(List<Integer> scores) {
        this.scores = scores
    }

    static int latest() {
        return scores.last()
    }

    static int personalBest() {
        return scores.max()
    }

    static def personalTopThree() {
        return this.scores.sort(false, {-it}).take(3)
    }
}