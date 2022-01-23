package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        return len(layers) * 2
    }
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodles int = 0
    var sauce float64 = 0
    for _, ingredient := range layers {
        if ingredient == "noodles" {
            noodles = noodles + 50
        } else if ingredient == "sauce" {
        	sauce = sauce + 0.2
        }
    }
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
    lastItem := friendsList[len(friendsList) - 1]
    newList := append(myList, lastItem)
    return newList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
    var newQuantities = make([]float64, len(quantities))
    var scalePortions float64 = float64(numberOfPortions) / 2
    for i := 0; i < len(quantities); i++ {
        newQuantities[i] = quantities[i] * scalePortions
    }
	return newQuantities
}
