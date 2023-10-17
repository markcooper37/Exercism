import java.lang.Character;

class SqueakyClean {
    static String clean(String identifier) {
        StringBuilder output = new StringBuilder();
        boolean previousDash = false;
        for (int i = 0; i < identifier.length(); i++) {
            char character = identifier.charAt(i);
            if (character == '-') {
                previousDash = true;
                continue;
            } else if (Character.isISOControl(character)) {
                output.append("CTRL");
            } else if (Character.isWhitespace(character)) {
                output.append('_');
            } else if (Character.isLetter(character) && (character < 'α' || character > 'ω')) {
                if (previousDash) {
                    output.append(Character.toUpperCase(character));
                } else {
                    output.append(character);
                }
            }
            previousDash = false;
        }
        return output.toString();
    }
}
