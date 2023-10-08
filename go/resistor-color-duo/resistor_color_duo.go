package resistorcolorduo

// Value returns the resistance value of a resistor with given colors.
func Value(colors []string) int {
	colorMap := map[string]int{"black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4, "green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9}
    return 10 * colorMap[colors[0]] + colorMap[colors[1]]
}
