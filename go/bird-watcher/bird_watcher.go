package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    count := 0
	for i := 0; i < len(birdsPerDay); i++ {
        count = count + birdsPerDay[i]
    }
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    count := 0
    firstDayOfWeek := 7 * (week - 1)
	for i := 0; i < 7; i++ {
        count = count + birdsPerDay[firstDayOfWeek + i]
    }
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i = i + 2 {
        birdsPerDay[i] = birdsPerDay[i] + 1
    }
	return birdsPerDay
}
