class Anagram {

    def subject

    Anagram(subject) {
        this.subject = subject
    }

    def find(candidates) {
        def orderedSubject = orderLetters(subject)
        return candidates.findAll { it.toLowerCase() != subject.toLowerCase() && orderLetters(it) == orderedSubject }
    }

    def private orderLetters(str) {
        return str.toLowerCase().split("").sort()
    }

}