#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    if _first_list == _second_list {
        Comparison::Equal
    } else if _first_list.len() < _second_list.len() && is_first_sublist(_first_list, _second_list)
    {
        Comparison::Sublist
    } else if _first_list.len() > _second_list.len() && is_first_sublist(_second_list, _first_list)
    {
        Comparison::Superlist
    } else {
        Comparison::Unequal
    }
}

pub fn is_first_sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> bool {
    (0..=_second_list.len() - _first_list.len()).any(|i| {
        _second_list
            .iter()
            .skip(i)
            .zip(_first_list.iter())
            .all(|x| x.0 == x.1)
    })
}
