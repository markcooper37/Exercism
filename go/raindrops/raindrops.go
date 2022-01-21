package raindrops

import "strconv"

func Convert(number int) string {
	ret := ""
    if number % 3 == 0 {
        ret = ret + "Pling"
    }
	if number % 5 == 0 {
        ret = ret + "Plang"
    }
	if number % 7 == 0 {
        ret = ret + "Plong"
    }
	if ret == "" {
        return strconv.Itoa(number)
    } else {
    	return ret
    }
}
