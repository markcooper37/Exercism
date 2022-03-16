#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut numbers: Vec<i32> = vec![];
    for input in inputs {
        match input {
            CalculatorInput::Add => {
                let first = numbers.pop();
                let second = numbers.pop();
                match second {
                    Some(_) => {numbers.push(second.unwrap()+first.unwrap());},
                    None => return None,
                }
            },
            CalculatorInput::Subtract => {
                let first = numbers.pop();
                let second = numbers.pop();
                match second {
                    Some(_) => {numbers.push(second.unwrap()-first.unwrap());},
                    None => return None,
                }
            },
            CalculatorInput::Multiply => {
                let first = numbers.pop();
                let second = numbers.pop();
                match second {
                    Some(_) => {numbers.push(second.unwrap()*first.unwrap());},
                    None => return None,
                }
            },
            CalculatorInput::Divide => {
                let first = numbers.pop();
                let second = numbers.pop();
                match second {
                    Some(_) => {numbers.push(second.unwrap()/first.unwrap());},
                    None => return None,
                }
            },
            CalculatorInput::Value(number) => {
                numbers.push(*number);
            },
        }
    }
    let result = numbers.pop();
    match result {
        None => None,
        Some(_) => {
            if numbers.len() != 0 {
                None
            } else {
                result
            }
        },
    }
}
