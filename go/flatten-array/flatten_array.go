package flatten

func Flatten(nested interface{}) []interface{} {
	newList := []interface{}{}
	something, ok := nested.([]interface{})
	if !ok {
		if nested != nil {
			newList = append(newList, nested.(int))
		}
		return newList
	}
	for _, element := range something {
		newList = append(newList, Flatten(element))
	}
	flattenedList := []interface{}{}
	for _, value := range newList {
		for _, element := range value.([]interface{}) {
			flattenedList = append(flattenedList, element.(int))
		}
	}
	return flattenedList
}
