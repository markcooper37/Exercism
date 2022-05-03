package zebra

type House struct {
	nationality string
	colour      string
	pet         string
	drink       string
	cigarettes  string
}

type Houses [5]House

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolvePuzzle() Solution {
	solutionsStepOne := []Houses{}
	for i := 0; i < 120; i++ {
		seed := i
		houses := Houses{}
		colours := []string{"red", "green", "yellow", "blue", "ivory"}
		for j := 5; j >= 1; j-- {
			houses[5-j].colour = colours[seed%j]
			colours = append(colours[:seed%j], colours[seed%j+1:]...)
			seed /= j
		}
		if houses.stepOneChecksPass() {
			solutionsStepOne = append(solutionsStepOne, houses)
		}
	}
	solutionsStepTwo := []Houses{}
	for i := 0; i < 120; i++ {
		for _, houses := range solutionsStepOne {
			seed := i
			nationalities := []string{"Englishman", "Spaniard", "Ukrainian", "Norwegian", "Japanese"}
			for j := 5; j >= 1; j-- {
				houses[5-j].nationality = nationalities[seed%j]
				nationalities = append(nationalities[:seed%j], nationalities[seed%j+1:]...)
				seed /= j
			}
			if houses.stepTwoChecksPass() {
				solutionsStepTwo = append(solutionsStepTwo, houses)
			}
		}
	}
	solutionsStepThree := []Houses{}
	for i := 0; i < 120; i++ {
		for _, houses := range solutionsStepTwo {
			seed := i
			pets := []string{"dog", "snails", "fox", "horse", "zebra"}
			for j := 5; j >= 1; j-- {
				houses[5-j].pet = pets[seed%j]
				pets = append(pets[:seed%j], pets[seed%j+1:]...)
				seed /= j
			}
			if houses.stepThreeChecksPass() {
				solutionsStepThree = append(solutionsStepThree, houses)
			}
		}
	}
	solutionsStepFour := []Houses{}
	for i := 0; i < 120; i++ {
		for _, houses := range solutionsStepThree {
			seed := i
			drinks := []string{"coffee", "tea", "milk", "orange juice", "water"}
			for j := 5; j >= 1; j-- {
				houses[5-j].drink = drinks[seed%j]
				drinks = append(drinks[:seed%j], drinks[seed%j+1:]...)
				seed /= j
			}
			if houses.stepFourChecksPass() {
				solutionsStepFour = append(solutionsStepFour, houses)
			}
		}
	}
	solutionsStepFive := []Houses{}
	for i := 0; i < 120; i++ {
		for _, houses := range solutionsStepFour {
			seed := i
			cigarettes := []string{"Old Gold", "Kool", "Chesterfield", "Lucky Strike", "Parliament"}
			for j := 5; j >= 1; j-- {
				houses[5-j].cigarettes = cigarettes[seed%j]
				cigarettes = append(cigarettes[:seed%j], cigarettes[seed%j+1:]...)
				seed /= j
			}
			if houses.stepFiveChecksPass() {
				solutionsStepFive = append(solutionsStepFive, houses)
			}
		}
	}
	completeHouses := solutionsStepFive[0]
	solution := Solution{}
	for i := 0; i <= 4; i++ {
		if completeHouses[i].pet == "zebra" {
			solution.OwnsZebra = completeHouses[i].nationality
		} else if completeHouses[i].drink == "water" {
			solution.DrinksWater = completeHouses[i].nationality
		}
	}
	return solution
}

func (h *Houses) stepOneChecksPass() bool {
	for i := 0; i <= 3; i++ {
		if h[i].colour == "ivory" && h[i+1].colour == "green" {
			return true
		}
	}
	return false
}

func (h *Houses) stepTwoChecksPass() bool {
	for i := 0; i <= 4; i++ {
		if h[i].nationality == "Englishman" && h[i].colour != "red" {
			return false
		}
		if h[i].nationality == "Norwegian" && (i == 0 || h[i-1].colour != "blue") && (i == 4 || h[i+1].colour != "blue") {
			return false
		}
	}
	return h[0].nationality == "Norwegian"
}

func (h *Houses) stepThreeChecksPass() bool {
	for i := 0; i <= 4; i++ {
		if h[i].nationality == "Spaniard" && h[i].pet != "dog" {
			return false
		}
	}
	return true
}

func (h *Houses) stepFourChecksPass() bool {
	for i := 0; i <= 4; i++ {
		if h[i].drink == "coffee" && h[i].colour != "green" {
			return false
		}
		if h[i].drink == "tea" && h[i].nationality != "Ukrainian" {
			return false
		}
	}
	return h[2].drink == "milk"
}

func (h *Houses) stepFiveChecksPass() bool {
	for i := 0; i <= 4; i++ {
		if h[i].cigarettes == "Old Gold" && h[i].pet != "snails" {
			return false
		}
		if h[i].cigarettes == "Kool" && h[i].colour != "yellow" {
			return false
		}
		if h[i].cigarettes == "Lucky Strike" && h[i].drink != "orange juice" {
			return false
		}
		if h[i].cigarettes == "Parliament" && h[i].nationality != "Japanese" {
			return false
		}
		if h[i].cigarettes == "Chesterfield" && (i == 0 || h[i-1].pet != "fox") && (i == 4 || h[i+1].pet != "fox") {
			return false
		}
		if h[i].cigarettes == "Kool" && (i == 0 || h[i-1].pet != "horse") && (i == 4 || h[i+1].pet != "horse") {
			return false
		}
	}
	return true
}
