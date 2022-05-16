use std::collections::HashMap;

pub fn primes_up_to(upper_bound: u64) -> Vec<u64> {
    let mut prime_map = HashMap::new();
    for i in 2..=upper_bound {
        if prime_map.get(&i) == None {
            for j in 2..=(upper_bound / i) {
                prime_map.insert(i * j, false);
            }
        }
    }
    (2..=upper_bound)
        .filter(|c| prime_map.get(c) == None)
        .collect::<Vec<u64>>()
}
