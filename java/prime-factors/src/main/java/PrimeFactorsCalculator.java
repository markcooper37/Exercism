import java.util.List;
import java.util.ArrayList;

class PrimeFactorsCalculator {

    List<Long> calculatePrimeFactorsOf(long number) {
        long factor = 2;
        List<Long> factors = new ArrayList<Long>();
        while (number > 1) {
            if (number % factor == 0) {
                factors.add(factor);
                number /= factor;
            } else {
                factor++;
            }
        }
        return factors;
    }

}