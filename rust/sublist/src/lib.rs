use std::cmp::Ordering;

#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    match _first_list.len().cmp(&_second_list.len()) {
        Ordering::Equal => {
            match _first_list == _second_list {
                true => Comparison::Equal,
                false => Comparison::Unequal
            }
        }
        Ordering::Less => {
            match is_first_sublist(_first_list, _second_list) {
                true => Comparison::Sublist,
                false => Comparison::Unequal
            }
        }
        Ordering::Greater => {
            match is_first_sublist(_second_list, _first_list) {
                true => Comparison::Superlist,
                false => Comparison::Unequal
            }
        }
    }
}

pub fn is_first_sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> bool {
    (0..=_second_list.len() - _first_list.len()).any(|i| _second_list
        .iter()
        .skip(i)
        .zip(_first_list.iter())
        .all(|x| x.0 == x.1))
}