package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * successRate / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(float64(productionRate) * successRate / 6000)
}

func CalculateCost(carsCount int) uint {
    tensOfCars := carsCount / 10
    remainderCars := carsCount % 10
    return uint(tensOfCars * 95000 + remainderCars * 10000)
}
