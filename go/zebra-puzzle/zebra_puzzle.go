package zebra

import "fmt"

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
	grids := []Grid{}
	for i := 0; i <= 4; i++ {
		for j := i + 1; j <= 5; j++ {
			grids = append(grids, Grid{rows: Category(i), columns: Category(j), grid: [5][5]int{}})
		}
	}
	for index := range grids {
		grids[index].InitialiseGrid()
	}
	gridsComplete := false
	for !gridsComplete {
		gridsComplete = true
		for index := range grids {
			grids[index].CompareGrids(grids)
			grids[index].AdditionalChecks()
			grids[index].RowColumnComplete()
			grids[index].FillSinglePossibilities()
			if !grids[index].IsComplete() {
				gridsComplete = false
			}
		}
		fmt.Println(grids)
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
	if g.rows == Position && g.columns == Colour {
		g.grid[0][1] = NoMatch
		g.grid[4][4] = NoMatch
	}
}

func (g *Grid) IsComplete() bool {
	for _, row := range g.grid {
		for _, position := range row {
			if position == Unknown {
				return false
			}
		}
	}
	return true
}

func (g *Grid) CompareGrids(grids []Grid) {
	for i := 0; i <= 5; i++ {
		if i == int(g.rows) || i == int(g.columns) {
			continue
		}
		firstGrid, secondGrid := Grid{}, Grid{}
		if i < int(g.rows) {
			firstGrid = grids[FindGridIndex(i, int(g.rows))]
		} else {
			firstGrid = grids[FindGridIndex(int(g.rows), i)]
		}
		if i < int(g.columns) {
			secondGrid = grids[FindGridIndex(i, int(g.columns))]
		} else {
			secondGrid = grids[FindGridIndex(int(g.columns), i)]
		}
		for rowIndex, row := range g.grid {
			for columnIndex := range row {
				if g.grid[rowIndex][columnIndex] != 0 {
					continue
				}
				for j := 0; j <= 4; j++ {
					isFirstMatch, isSecondMatch := 0, 0
					if i < int(g.rows) {
						isFirstMatch = firstGrid.grid[j][rowIndex]
					} else {
						isFirstMatch = firstGrid.grid[rowIndex][j]
					}
					if i < int(g.columns) {
						isSecondMatch = secondGrid.grid[j][columnIndex]
					} else {
						isSecondMatch = secondGrid.grid[columnIndex][j]
					}
					if isFirstMatch == Match && isSecondMatch == NoMatch || isFirstMatch == NoMatch && isSecondMatch == Match {
						g.grid[rowIndex][columnIndex] = NoMatch
					} else if isFirstMatch == Match && isSecondMatch == Match {
						g.grid[rowIndex][columnIndex] = Match
					}
				}
			}
		}
	}
}

func (g *Grid) AdditionalChecks() {
	// checking green house is on the right of the ivory house
	if g.rows == Position && g.columns == Colour {
		for i := 0; i <= 4; i++ {
			if g.grid[i][1] == Match {
				g.grid[i-1][4] = Match
			} else if g.grid[i][1] == NoMatch && i > 0 {
				g.grid[i-1][4] = NoMatch
			}
			if g.grid[i][4] == Match {
				g.grid[i+1][1] = Match
			} else if g.grid[i][4] == NoMatch && i < 4 {
				g.grid[i+1][1] = NoMatch
			}
		}
	}
	// checking Chesterfield is next to fox

	// checking Kool is next to horse

	// checking Norwegian is next to blue
}

func (g *Grid) RowColumnComplete() {
	for rowIndex, row := range g.grid {
		for columnIndex := range row {
			if g.RowComplete(rowIndex) || g.ColumnComplete(columnIndex) {
				g.grid[rowIndex][columnIndex] = NoMatch
			}
		}
	}
}

func (g *Grid) RowComplete(rowIndex int) bool {
	for _, position := range g.grid[rowIndex] {
		if position == Match {
			return true
		}
	}
	return false
}

func (g *Grid) ColumnComplete(columnIndex int) bool {
	for _, row := range g.grid {
		if row[columnIndex] == Match {
			return true
		}
	}
	return false
}

func (g *Grid) FillSinglePossibilities() {
	for rowIndex := range g.grid {
		unknownIndex, oneUnknown := g.OnePossibilityRow(rowIndex)
		if oneUnknown {
			g.grid[rowIndex][unknownIndex] = Match
		}
	}
	for columnIndex := range g.grid[0] {
		unknownIndex, oneUnknown := g.OnePossibilityColumn(columnIndex)
		if oneUnknown {
			g.grid[unknownIndex][columnIndex] = Match
		}
	}
}

func (g *Grid) OnePossibilityRow(rowIndex int) (int, bool) {
	oneUnknown := false
	unknownIndex := -1
	for index, value := range g.grid[rowIndex] {
		if value == Match {
			return -1, false
		} else if value == Unknown && !oneUnknown {
			unknownIndex = index
			oneUnknown = true
		} else if value == Unknown && oneUnknown {
			return -1, false
		}
	}
	return unknownIndex, oneUnknown
}

func (g *Grid) OnePossibilityColumn(columnIndex int) (int, bool) {
	oneUnknown := false
	unknownIndex := -1
	for index, row := range g.grid {
		if row[columnIndex] == Match {
			return -1, false
		} else if row[columnIndex] == Unknown && !oneUnknown {
			unknownIndex = index
			oneUnknown = true
		} else if row[columnIndex] == Unknown && oneUnknown {
			return -1, false
		}
	}
	return unknownIndex, oneUnknown
}

func FindGridIndex(row, column int) int {
	if row == 0 {
		return column
	} else if row == 1 {
		return row + column + 2
	} else if row == 2 {
		return row + column + 4
	} else if row == 3 {
		return row + column + 5
	} else {
		return row + column + 5
	}
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
