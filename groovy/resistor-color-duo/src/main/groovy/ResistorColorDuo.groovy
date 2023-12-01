class ResistorColorDuo {

    static colors = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']

    static int value(List<String> colorsInput) {
        return colors.indexOf(colorsInput[0]) * 10 + colors.indexOf(colorsInput[1])
    }
}