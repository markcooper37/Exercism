// Package triangle contains a function to determine a triangle's kind from its side lengths.
package triangle

type Kind string

const (
    NaT Kind = "NaT"
    Equ Kind = "Equ"
    Iso Kind = "Iso"
    Sca Kind = "Sca"
)

// KindFromSides determines a triangle's kind from the length of its sides.
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
        return NaT
    } else if a + b <= c || a + c <= b || b + c <= a {
        return NaT
    } else if a == b && b == c {
    	return Equ
    } else if a == b || b == c || a == c {
    	return Iso
    } else {
    	return Sca
    }
	
}
