use rand::Rng;

pub fn encode(key: &str, s: &str) -> Option<String> {
    if !key.chars().all(|c| c >= 'a' && c <= 'z')
        || !s.chars().all(|c| c >= 'a' && c <= 'z')
        || key.is_empty()
    {
        return None;
    }
    let key = key.chars().collect::<Vec<char>>();
    let s = s
        .chars()
        .enumerate()
        .map(|c| (((c.1 as u8 + key[c.0 % key.len()] as u8 - 194) % 26) + 97) as char)
        .collect::<String>();
    Some(s)
}

pub fn decode(key: &str, s: &str) -> Option<String> {
    if !key.chars().all(|c| c >= 'a' && c <= 'z')
        || !s.chars().all(|c| c >= 'a' && c <= 'z')
        || key.is_empty()
    {
        return None;
    }
    let key = key.chars().collect::<Vec<char>>();
    let s = s
        .chars()
        .enumerate()
        .map(|c| (((c.1 as u8 + 26 - key[c.0 % key.len()] as u8) % 26) + 97) as char)
        .collect::<String>();
    Some(s)
}

pub fn encode_random(s: &str) -> (String, String) {
    let chars: &[u8] = b"abcdefghijklmnopqrstuvwxyz";
    let key_length: usize = 100;
    let mut rng = rand::thread_rng();

    let key: String = (0..key_length)
        .map(|_| {
            let idx = rng.gen_range(0..chars.len());
            chars[idx] as char
        })
        .collect();

    let encoded_value = encode(&key, s).unwrap();

    (key, encoded_value)
}
