import java.util.List;
import java.util.ArrayList;

class Series {

    private String digits;
    
    Series(String string) {
        this.digits = string;
    }

    List<String> slices(int num) {
        if (num > this.digits.length()) {
            throw new IllegalArgumentException("Slice size is too big.");
        } else if (num < 1) {
            throw new IllegalArgumentException("Slice size is too small.");
        }
        List<String> substrings = new ArrayList();
        for (int i = 0; i <= this.digits.length() - num; i++) {
            substrings.add(this.digits.substring(i, i + num));
        }
        return substrings;
    }
}
