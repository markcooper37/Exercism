
class Matrix {

    private String matrixAsString;

    Matrix(String matrixAsString) {
        this.matrixAsString = matrixAsString;
    }

    int[] getRow(int rowNumber) {
        String[] rows = matrixAsString.split("\n");
        String[] requiredRow = rows[rowNumber-1].split(" ");
        int[] output = new int[requiredRow.length];
        for (int i = 0; i < requiredRow.length; i++) {
            output[i] = Integer.valueOf(requiredRow[i]);
        }
        return output;
    }

    int[] getColumn(int columnNumber) {
        String[] rows = matrixAsString.split("\n");
        int[] output = new int[rows.length];
        for (int i = 0; i < rows.length; i++) {
            String[] requiredRow = rows[i].split(" ");
            output[i] = Integer.valueOf(requiredRow[columnNumber-1]);
        }
        return output;
    }
}
