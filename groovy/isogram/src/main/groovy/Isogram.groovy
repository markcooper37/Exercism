class Isogram {
    static boolean isIsogram(String phrase) {
        def letters = phrase.replaceAll("[^A-Za-z]", "").toLowerCase()
        return letters.toSet().size() == letters.length()
    }
}