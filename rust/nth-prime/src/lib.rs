pub fn nth(n: u32) -> u32 {
    let mut primes: Vec<u32> = vec![];
    let mut test_number: u32 = 2;
    let result: u32 = loop {
        let mut is_prime = true;
        for prime in &primes {
            if test_number % prime == 0 {
                is_prime = false;
                break;
            }
        }
        if is_prime {
            primes.push(test_number);
        }
        if (primes.len() as u32) == n + 1 {
            break test_number;
        }
        test_number += 1;
    };
    result
}
