class RunLengthEncoding {

    static encode(input) {
        return input.findAll(/(.)\1{0,}/).collect {"${it.length() > 1 ? it.length() : ""}${it[0]}"} .join("")
    }

    static decode(input) {
        return input.findAll(/\d*\D/).collect {
            "${it.substring(it.length()-1) * (it.length() > 1 ? it.substring(0, it.length()-1).toInteger() : 1)}"
        } .join("")
    }
}