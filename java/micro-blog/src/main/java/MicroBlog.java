class MicroBlog {
    public String truncate(String input) {
        int cutoff = (int) input.codePoints().count();
        if (cutoff > 5) {
            cutoff = 5;
        }
        return input.substring(0, input.offsetByCodePoints(0, cutoff));
    }
}
