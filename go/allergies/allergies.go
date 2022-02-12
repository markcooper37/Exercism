package allergies

var stringAllergies = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(allergies uint) []string {
	validAllergies := []string{}
	for i := 0; i < 8; i++ {
		if (allergies >> i) & 1 == 1 {
			validAllergies = append(validAllergies, stringAllergies[i])
		}
	}
	return validAllergies
}

func AllergicTo(allergies uint, allergen string) bool {
	for i, stringAllergy := range stringAllergies {
		if allergen == stringAllergy {
			if (allergies >> i) & 1 == 1 {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
