class SecretHandshake {

    def static actions = ["wink", "double blink", "close your eyes", "jump"]

    static List<String> commands(int number) {
        def output = []
        for (def i = 0; i <= 3; i++) {
            if (number & 1 << i) {
                output.add(actions[i])
            }
        }
        if (number & 1 << 4) {
            output = output.reverse()
        }
        return output
    }
}
