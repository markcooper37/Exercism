package bookstore

import "sort"

func Cost(books []int) int {
	quantities := make([]int, 5)
	for _, book := range books {
		quantities[book-1]++
	}
	sort.Ints(quantities)
	counts := make([]int, 5)
	for quantities[4] > 0 {
		count := 0
		for i := 0; i <= 4; i++ {
			if quantities[i] > 0 {
				count++
				quantities[i]--
			}
		}
		counts[count-1]++
	}
	for counts[2] > 0 && counts[4] > 0 {
		counts[2]--
		counts[4]--
		counts[3] += 2
	}
	return 800 * counts[0] + 1520 * counts[1] + 2160 * counts[2] + 2560 * counts[3] + 3000 * counts[4]
}
