import java.util.List;
import java.util.ArrayList;

class PrimeCalculator {

    int nth(int nth) {
        if (nth < 1) {
            throw new IllegalArgumentException("Input must be a positive integer.");
        }
        List<Integer> primes = new ArrayList<>();
        primes.add(2);
        int currentInteger = 3;
        while (primes.size() < nth) {
            boolean isPrime = true;
            for (int prime: primes) {
                if (currentInteger % prime == 0) {
                    isPrime = false;
                    break;
                }
            }
            if (isPrime) {
                primes.add(currentInteger);
            }
            currentInteger++;
        }
        return primes.get(nth - 1);
    }

}
