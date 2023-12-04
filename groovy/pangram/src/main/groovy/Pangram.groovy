class Pangram {

    static boolean isPangram(String sentence) {
        return sentence.replaceAll("[^A-Za-z]", "").toLowerCase().toSet().size() == 26
    }

}