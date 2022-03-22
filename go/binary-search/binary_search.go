package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 {
		return -1
	}
	pos := len(list) / 2
	if list[pos] == key {
		return pos
	} else if list[pos] < key {
		finalPos :=  SearchInts(list[pos+1:], key)
		if finalPos == -1 {
			return -1
		} else {
			return finalPos + pos + 1
		}
	} else {
		return SearchInts(list[0:pos], key)
	}
}
