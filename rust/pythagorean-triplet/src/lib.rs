use std::collections::HashSet;

pub fn find(sum: u32) -> HashSet<[u32; 3]> {
    let mut triplets: HashSet<[u32; 3]> = HashSet::new();
    for a in 1..=sum/3 {
        for b in a..=sum/2 {
            for c in b..=sum {
                if a.pow(2) + b.pow(2) == c.pow(2) && a + b + c == sum {
                    triplets.insert([a, b, c]);
                }
            }
        }
    }
    triplets
}
