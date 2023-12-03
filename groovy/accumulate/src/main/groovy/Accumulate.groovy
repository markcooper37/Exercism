class Accumulate {

    static accumulate(Collection collection, Closure func) {
        return collection.collect(func)
    }

}