pub fn map<S, T, F: FnMut(S) -> T>(input: Vec<S>, mut function: F) -> Vec<T> {
    let mut output = vec![];
    for value in input {
        output.push(function(value));
    }
    output
}
