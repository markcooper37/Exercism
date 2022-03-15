/// Creates an empty vector.
pub fn create_empty() -> Vec<u8> {
    vec![]
}

/// Create a buffer of `count` zeroes.
/// Applications often use buffers when serializing data to send over the network.
pub fn create_buffer(count: usize) -> Vec<u8> {
    vec![0; count]
}

/// Creates a vector containing the first five elements of the Fibonacci sequence.
pub fn fibonacci() -> Vec<u8> {
    vec![1, 1, 2, 3, 5]
}
