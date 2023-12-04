class ETL {

    static transform(input) {
        def output = [:]
        input.each { k, v -> for (val in v) { output[val.toLowerCase()] = k as int } }
        return output
    }

}