class ResistorColorTrio {

    static colors = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']

    static String label(List<String> colorsInput) {
        def ohms = (colors.indexOf(colorsInput[0]) * 10 + colors.indexOf(colorsInput[1])) * (10 ** colors.indexOf(colorsInput[2]))
        if (ohms < 1000) {
            "${ohms} ohms"
        } else {
            "${ohms.intdiv(1000)} kiloohms"
        }
    }

}