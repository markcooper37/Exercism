class Raindrops {

    def convert(Integer num) {
        def sound = ""
        if (num % 3 == 0) {
            sound += "Pling"
        }
        if (num % 5 == 0) {
            sound += "Plang"
        }
        if (num % 7 == 0) {
            sound += "Plong"
        }
        if (sound == "") {
            sound = "$num"
        }
        return sound
    }

}
