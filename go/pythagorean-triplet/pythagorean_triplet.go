package pythagorean

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	triplets := []Triplet{}
    for a := min; a <= max; a++ {
        for b := a; b <= max; b++ {
            for c := b; c <= max; c++ {
                if a * a + b * b == c * c {
                    triplets = append(triplets, [3]int{a, b, c})
                }
            }
        }
    }
	return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
    possibleTriplets := Range(1, p)
    triplets := []Triplet{}
    for _, triplet := range possibleTriplets {
        if triplet[0] + triplet[1] + triplet[2] == p {
            triplets = append(triplets, triplet)
        }
    }
	return triplets
}
