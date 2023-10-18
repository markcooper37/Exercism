class QueenAttackCalculator {

    private final int q1Row;
    private final int q1Col;
    private final int q2Row;
    private final int q2Col;

    QueenAttackCalculator(Queen queen1, Queen queen2) {
        if (queen1 == null || queen2 == null) {
            throw new IllegalArgumentException("You must supply valid positions for both Queens.");
        } else if (queen1.row == queen2.row && queen1.column == queen2.column) {
            throw new IllegalArgumentException("Queens cannot occupy the same position.");
        }
        this.q1Row = queen1.row;
        this.q1Col = queen1.column;
        this.q2Row = queen2.row;
        this.q2Col = queen2.column;
    }

    boolean canQueensAttackOneAnother() {
        return q1Row == q2Row
            || q1Col == q2Col
            || q1Row + q1Col == q2Row + q2Col
            || q1Row + q2Col == q2Row + q1Col;
    }

}