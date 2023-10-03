package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for index, resistorColor := range Colors() {
        if color == resistorColor {
            return index
        }
    }
	return -1
}
