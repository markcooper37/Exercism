convert(N, Numeral) :- get_units(N, Units), get_tens(N, Tens), get_hundreds(N, Hundreds), get_thousands(N, Thousands), atomics_to_string([Thousands, Hundreds, Tens, Units], Numeral).

get_units(N, Units) :- Digit is mod(N, 10), nth0(Digit, ["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"], Units).

get_tens(N, Tens) :- Digit is mod(div(N, 10), 10), nth0(Digit, ["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"], Tens).

get_hundreds(N, Hundreds) :- Digit is mod(div(N, 100), 10), nth0(Digit, ["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"], Hundreds).

get_thousands(N, Thousands) :- Digit is mod(div(N, 1000), 10), nth0(Digit, ["", "M", "MM", "MMM"], Thousands).
