package resistorcolortrio

import (
    "math"
    "strconv"
)

// Label describes the resistance value given the colors of a resistor.
func Label(colors []string) string {
	colorMap := map[string]int{"black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4, "green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9}
    ohms := (10 * colorMap[colors[0]] + colorMap[colors[1]]) * int(math.Pow10(colorMap[colors[2]]))
    if ohms >= 1000000000 {
        return strconv.Itoa(ohms/1000000000) + " gigaohms"
    } else if ohms >= 1000000 {
    	return strconv.Itoa(ohms/1000000) + " megaohms"
    } else if ohms >= 1000 {
    	return strconv.Itoa(ohms/1000) + " kiloohms"
    } else {
    	return strconv.Itoa(ohms) + " ohms"
    }
}
