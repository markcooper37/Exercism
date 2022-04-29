package cards

// FavoriteCards returns a slice containing 2, 6 and 9.
func FavoriteCards() []int {
    return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position, or -1 if the position doesn't exist.
func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
        return -1
    } else {
    	return slice[index]
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
        slice = append(slice, value)
        return slice
    } else {
    	slice[index] = value
        return slice
    }
}

// PrependItems adds items to the front of the slice.
func PrependItems(slice []int, value ...int) []int {
    return append(value, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index > len(slice) || index < 0 {
        return slice
    } else {
    	sliceStart := slice[:index]
        sliceEnd := slice[index+1:]
        newSlice := append(sliceStart, sliceEnd...)
        return newSlice
    }
}
