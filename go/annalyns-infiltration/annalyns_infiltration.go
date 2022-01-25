package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
        return false
    }
	return true
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
        return true
    }
	return false
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if prisonerIsAwake && !archerIsAwake {
        return true
    }
	return false
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if !archerIsAwake {
        if petDogIsPresent || (prisonerIsAwake && !knightIsAwake) {
            return true
        }
    }
	return false
}
