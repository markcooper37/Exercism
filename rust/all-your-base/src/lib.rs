#[derive(Debug, PartialEq)]
pub enum Error {
    InvalidInputBase,
    InvalidOutputBase,
    InvalidDigit(u32),
}

///
/// Convert a number between two bases.
///
/// A number is any slice of digits.
/// A digit is any unsigned integer (e.g. u8, u16, u32, u64, or usize).
/// Bases are specified as unsigned integers.
///
/// Return an `Err(.)` if the conversion is impossible.
/// The tests do not test for specific values inside the `Err(.)`.
///
///
/// You are allowed to change the function signature as long as all test still pass.
///
///
/// Example:
/// Input
///   number: &[4, 2]
///   from_base: 10
///   to_base: 2
/// Result
///   Ok(vec![1, 0, 1, 0, 1, 0])
///
/// The example corresponds to converting the number 42 from decimal
/// which is equivalent to 101010 in binary.
///
///
/// Notes:
///  * The empty slice ( "[]" ) is equal to the number 0.
///  * Never output leading 0 digits, unless the input number is 0, in which the output must be `[0]`.
///    However, your function must be able to process input with leading 0 digits.
///
pub fn convert(number: &[u32], from_base: u32, to_base: u32) -> Result<Vec<u32>, Error> {
    if from_base <= 1 {
        Err(Error::InvalidInputBase)
    } else if to_base <= 1 {
        Err(Error::InvalidOutputBase)
    } else {
        for value in number {
            if *value >= from_base {
                return Err(Error::InvalidDigit(*value));
            }
        }
        let mut number: u32 = number
            .iter()
            .rev()
            .enumerate()
            .map(|n| n.1 * from_base.pow(n.0 as u32))
            .sum();
        let mut output: Vec<u32> = vec![];
        let mut index = 0;
        while number > 0 {
            let value_to_subtract = number % to_base.pow(index + 1);
            output.push(value_to_subtract / to_base.pow(index));
            number -= value_to_subtract;
            index += 1;
        }
        if output.len() == 0 {
            output.push(0);
        }
        Ok(output.into_iter().rev().collect())
    }
}
