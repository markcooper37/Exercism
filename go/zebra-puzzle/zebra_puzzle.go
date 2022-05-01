package zebra

type Grid struct {
	rows Category
	columns Category
	grid [5][5]int
}

type Category int

const(
	Position Category = iota
	Nationality
	Colour
	Pet
	Drink
	Cigarettes
)

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolvePuzzle() Solution {
	panic("Please implement the SolvePuzzle function")
}

// Englishman, Spaniard, Ukranian, Norwegian, Japanese
// red, green, yellow, blue, ivory
// dog, snails, fox, horse, zebra
// coffee, tea, milk, orange juice, water
// Old Gold, Kool, Chesterfield, Lucky Strike, Parliament
// first, second, third, fourth, fifth

// Create grids which pair each of the above categories where there is exactly one tick in each row and column
// Fill in what is given initially
// Iteratively fill in more information, checking the initial requirements about relative house positions that could not be satisfied at the start
// Use the fact that if a has b and b does not have c, a cannot have c
// For example, if the Englishman lives in the red house and the red house is not the first house,
// the Englishman does not live in the first house
// Use also the fact that is a has b and b has c, a has c
// For example, if the owner of the house with a pet dog drinks coffee and the homeowner that drinks coffee smokes Old Gold,
// the owner of the house with a pet dog must smoke Old Gold
// Finish when all grids are complete

// For an empty cell in a grid, check if the cell can be filled either with a tick or a cross
// If a tick, fill the rest of the row and column with crosses
// Check if any row/column has a single space with no cross and fill in if necessary
