class Queen {

    def row
    def column

    Queen(int row, int column) {
        if (row < 0 || row > 7 || column < 0 || column > 7) {
            throw new Exception("row and column must be between 0 and 7")
        }
        this.row = row
        this.column = column
    }

}
