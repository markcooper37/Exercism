pub fn is_armstrong_number(num: u32) -> bool {
    let mut temp_num = num;
    let power = num.to_string().len() as u32;
    let mut sum = 0;
    while temp_num != 0 {
        sum += u32::pow(temp_num % 10, power);
        temp_num /= 10;
    }
    return sum == num;
}
