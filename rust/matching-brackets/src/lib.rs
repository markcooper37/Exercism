use std::collections::HashMap;

pub fn brackets_are_balanced(string: &str) -> bool {
    let mut stack = vec![];
    let bracket_matches: HashMap<char, char> = HashMap::from([('{', '}'), ('[', ']'), ('(', ')')]);
    for bracket in string.chars() {
        match bracket {
            '{' | '[' | '(' => stack.push(bracket),
            '}' | ']' | ')' => match stack.pop() {
                None => return false,
                Some(open_bracket) => {
                    if *bracket_matches.get(&open_bracket).unwrap() != bracket {
                        return false;
                    }
                }
            },
            _ => (),
        }
    }
    stack.is_empty()
}
