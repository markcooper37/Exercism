pub fn collatz(n: u64) -> Option<u64> {
    if n == 0 {
        return None;
    }
    let mut m = n;
    let mut steps: u64 = 0;
    loop {
        if m == 1 {
            break;
        }
        if m % 2 == 0 {
            m = m / 2;
        } else if (u64::MAX - 1) / 3 < m {
            return None;
        } else {
            m = 3 * m + 1;
        }
        steps += 1;
    }
    Some(steps)
}
