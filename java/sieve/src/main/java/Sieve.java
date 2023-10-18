import java.util.List;
import java.util.Set;
import java.util.ArrayList;
import java.util.HashSet;

class Sieve {

    private final int maxPrime;
    
    Sieve(int maxPrime) {
        this.maxPrime = maxPrime;
    }

    List<Integer> getPrimes() {
        Set<Integer> composites = new HashSet<Integer>();
        List<Integer> primes = new ArrayList<Integer>();
        for (int i = 2; i <= this.maxPrime; i++) {
            if (!composites.contains(i)) {
                primes.add(i);
                for (int j = 2 * i; j <= this.maxPrime; j += i) {
                    composites.add(j);
                }
            }
        }
        return primes;
    }
}
