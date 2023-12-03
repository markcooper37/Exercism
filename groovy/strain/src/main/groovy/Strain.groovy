class Strain {

    static Collection keep(Collection collection, Closure predicate) {
        def output = []
        collection.each {
            if (predicate(it)) {
                output.add(it)
            }
        }
        return output
    }

    static Collection discard(Collection collection, Closure predicate) {
        def output = []
        collection.each {
            if (!predicate(it)) {
                output.add(it)
            }
        }
        return output
    }
}