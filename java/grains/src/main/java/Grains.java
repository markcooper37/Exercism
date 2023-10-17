import java.math.BigInteger;

class Grains {

    BigInteger grainsOnSquare(final int square) {
        if (square < 1 || square > 64) {
            throw new IllegalArgumentException("square must be between 1 and 64");
        }
        return new BigInteger("0").setBit(square - 1);
    }

    BigInteger grainsOnBoard() {
        BigInteger total = new BigInteger("0");
        for (int i = 1; i <= 64; i++) {
            total = total.add(grainsOnSquare(i));
        }
        return total;
    }

}
