package zebra

type Grid struct {
	rows    Category
	columns Category
	grid    [5][5]int
}

type Category int

const (
	Position Category = iota
	Nationality
	Colour
	Pet
	Drink
	Cigarettes
)

const (
	Unknown int = iota
	NoMatch
	Match
)

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolvePuzzle() Solution {
	grids := [5][6]Grid{}
	for i := 0; i <= 4; i++ {
		for j := i; j <= 5; j++ {
			grids[i][j] = Grid{rows: Category(i), columns: Category(j), grid: [5][5]int{}}
		}
	}
	for rowIndex, row := range grids {
		for columnIndex := range row {
			grids[rowIndex][columnIndex].InitialiseGrid()
		}
	}
	panic("Please implement the SolvePuzzle function")
}

func (g *Grid) InitialiseGrid() {
	if g == nil {
		return
	}
	if g.rows == Nationality && g.columns == Colour {
		g.grid[0][0] = Match
		g.grid[3][3] = NoMatch
	}
	if g.rows == Nationality && g.columns == Pet {
		g.grid[1][0] = Match
	}
	if g.rows == Colour && g.columns == Drink {
		g.grid[1][0] = Match
	}
	if g.rows == Nationality && g.columns == Drink {
		g.grid[2][1] = Match
	}
	if g.rows == Pet && g.columns == Cigarettes {
		g.grid[1][0] = Match
		g.grid[2][2] = NoMatch
		g.grid[3][1] = NoMatch
	}
	if g.rows == Colour && g.columns == Cigarettes {
		g.grid[2][1] = Match
	}
	if g.rows == Position && g.columns == Drink {
		g.grid[2][2] = Match
	}
	if g.rows == Position && g.columns == Nationality {
		g.grid[0][2] = Match
	}
	if g.rows == Drink && g.columns == Cigarettes {
		g.grid[3][3] = Match
	}
	if g.rows == Nationality && g.columns == Cigarettes {
		g.grid[2][2] = Match
	}
}

func (g *Grid) FillGrid() {

}

// Order in grids:
// Position: first, second, third, fourth, fifth
// Nationality: Englishman, Spaniard, Ukranian, Norwegian, Japanese
// Colour: red, green, yellow, blue, ivory
// Pet: dog, snails, fox, horse, zebra
// Drink: coffee, tea, milk, orange juice, water
// Cigarettes: Old Gold, Kool, Chesterfield, Lucky Strike, Parliament

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
