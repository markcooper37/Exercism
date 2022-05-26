use rand::Rng;

pub fn encode(key: &str, s: &str) -> Option<String> {
    if !key.chars().all(|c| c >= 'a' && c <= 'z')
        || !s.chars().all(|c| c >= 'a' && c <= 'z')
        || key.is_empty()
    {
        None
    } else {
        let key = key.chars().collect::<Vec<char>>();
        Some(
            s.chars()
                .enumerate()
                .map(|c| (((c.1 as u8 + key[c.0 % key.len()] as u8 - 194) % 26) + 97) as char)
                .collect::<String>(),
        )
    }
}

pub fn decode(key: &str, s: &str) -> Option<String> {
    if !key.chars().all(|c| c >= 'a' && c <= 'z')
        || !s.chars().all(|c| c >= 'a' && c <= 'z')
        || key.is_empty()
    {
        None
    } else {
        let key = key.chars().collect::<Vec<char>>();
        Some(
            s.chars()
                .enumerate()
                .map(|c| (((c.1 as u8 + 26 - key[c.0 % key.len()] as u8) % 26) + 97) as char)
                .collect::<String>(),
        )
    }
}

pub fn encode_random(s: &str) -> (String, String) {
    const CHARS: &[u8] = b"abcdefghijklmnopqrstuvwxyz";
    const KEY_LENGTH: usize = 100;
    let mut rng = rand::thread_rng();
    let key: String = (0..KEY_LENGTH)
        .map(|_| CHARS[rng.gen_range(0..CHARS.len())] as char)
        .collect();
    let encoded_value = encode(&key, s).unwrap();
    (key, encoded_value)
}
