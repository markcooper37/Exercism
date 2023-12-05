class MatchingBrackets {

    def static pairs = [']': '[', '}': '{', ')': '(']

    static isPaired(value) {
        def stack = []
        for (ch in value) {
            if (ch in ['[', '{', '(']) {
                stack.add(ch)
            } else if (ch in [']', '}', ')']) {
                if (stack.size() == 0 || stack.last() != pairs[ch]) {
                    return false
                } else {
                    stack.removeLast()
                }
            }
        }
        return stack.size() == 0
    }
}