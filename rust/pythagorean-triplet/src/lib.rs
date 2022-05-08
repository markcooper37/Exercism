use std::collections::HashSet;

pub fn find(sum: u32) -> HashSet<[u32; 3]> {
    let mut triplets: HashSet<[u32; 3]> = HashSet::new();
    for a in 1..=sum / 3 {
        for b in a..=sum / 2 {
            if a.pow(2) + b.pow(2) == (sum - a - b).pow(2) {
                triplets.insert([a, b, sum - a - b]);
            }
        }
    }
    triplets
}
