pub fn factors(n: u64) -> Vec<u64> {
    let mut factors: Vec<u64> = vec![];
    let mut n = n;
    let mut divisor = 2;
    while n > 1 {
        if n % divisor == 0 {
            n = n / divisor;
            factors.push(divisor)
        } else {
            divisor += 1;
        }
    }
    return factors;
}
