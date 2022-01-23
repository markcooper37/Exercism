package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        return len(layers) * 2
    }
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    for _, ingredient := range layers {
        if ingredient == "noodles" {
            noodles += 50
        } else if ingredient == "sauce" {
        	sauce += 0.2
        }
    }
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
    myList[len(myList)-1] = friendsList[len(friendsList) - 1]
    return myList
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
    var newQuantities = make([]float64, len(quantities))
    scalePortions := float64(numberOfPortions)/2
    for i := 0; i < len(quantities); i++ {
        newQuantities[i] = quantities[i] * scalePortions
    }
	return newQuantities
}
