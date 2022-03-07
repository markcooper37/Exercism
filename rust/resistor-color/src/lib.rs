use enum_iterator::IntoEnumIterator;
use int_enum::IntEnum;

#[repr(usize)]
#[derive(Clone, Copy, Debug, PartialEq, IntoEnumIterator, IntEnum)]

pub enum ResistorColor {
    Black = 0,
    Blue = 6,
    Brown = 1,
    Green = 5,
    Grey = 8,
    Orange = 3,
    Red = 2,
    Violet = 7,
    White = 9,
    Yellow = 4,
}

pub fn color_to_value(_color: ResistorColor) -> usize {
    _color as usize
}

pub fn enum_to_string(_color: ResistorColor) -> String {
    match _color {
        ResistorColor::Black => String::from("Black"),
        ResistorColor::Blue => String::from("Blue"),
        ResistorColor::Brown => String::from("Brown"),
        ResistorColor::Green => String::from("Green"),
        ResistorColor::Grey => String::from("Grey"),
        ResistorColor::Orange => String::from("Orange"),
        ResistorColor::Red => String::from("Red"),
        ResistorColor::Violet => String::from("Violet"),
        ResistorColor::White => String::from("White"),
        ResistorColor::Yellow => String::from("Yellow"),
    }
}

pub fn value_to_color_string(value: usize) -> String {
    for color in ResistorColor::into_enum_iter() {
        if color_to_value(color) == value {
            return enum_to_string(color);
        }
    };
    String::from("value out of range")
}

pub fn colors() -> Vec<ResistorColor> {
    let mut vec = Vec::with_capacity(10);
    for i in 0..10 {
        let res = ResistorColor::from_int(i).expect("something went wrong");
        vec.push(res);
    };
    vec
}
