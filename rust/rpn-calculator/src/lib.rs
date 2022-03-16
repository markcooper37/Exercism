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
            CalculatorInput::Value(number) => {
                numbers.push(*number);
            }
            _ => {
                if numbers.len() < 2 {
                    return None;
                }
                
                let second = numbers.pop().unwrap();
                let first = numbers.pop().unwrap();

                match input {
                    CalculatorInput::Add => numbers.push(first + second),
                    CalculatorInput::Subtract => numbers.push(first - second),
                    CalculatorInput::Multiply => numbers.push(first * second),
                    CalculatorInput::Divide => numbers.push(first / second),
                    _ => {}
                }
            }
        }
    }

    if numbers.len() != 1 {
        return None;
    }

    numbers.pop()
}
