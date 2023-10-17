public class Hamming {
    private final int hammingDistance;
    
    public Hamming(String leftStrand, String rightStrand) {
        if (leftStrand.length() == 0 && rightStrand.length() != 0) {
            throw new IllegalArgumentException("left strand must not be empty.");
        } else if (rightStrand.length() == 0 && leftStrand.length() != 0) {
            throw new IllegalArgumentException("right strand must not be empty.");
        } else if (leftStrand.length() != rightStrand.length()) {
            throw new IllegalArgumentException("leftStrand and rightStrand must be of equal length.");
        }
        int distance = 0;
        for (int i = 0; i < leftStrand.length(); i++) {
            if (leftStrand.charAt(i) != rightStrand.charAt(i)) {
                distance++;
            }
        }
        this.hammingDistance = distance;
    }

    public int getHammingDistance() {
        return this.hammingDistance;
    }
}
