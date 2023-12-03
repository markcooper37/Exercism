class ReverseString {

    static reverse(String value) {
        def output = ""
        for (character in value) {
            output = character + output
        }
        return output
    }

}