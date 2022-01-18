package sublist

type Relation string

const (
    equal Relation = "equal"
    sublist Relation = "sublist"
    superlist Relation = "superlist"
    unequal Relation = "unequal"
)

func Sublist(l1, l2 []int) Relation {
    sublistBool := false
    superlistBool := false
    if len(l1) == 0 {
        sublistBool = true
    }
	if len(l2) == 0 {
        superlistBool = true
    }
	if len(l1) <= len(l2) {
        for i := 0; i <= len(l2) - len(l1); i++ {
            for j := 0; j < len(l1); j++ {
                if l1[j] != l2[i + j] {
                    break
                }
            	if j == len(l1) - 1 {
                    sublistBool = true
                }
            }
        }
    }
    if len(l1) >= len(l2) {
        for i := 0; i <= len(l1) - len(l2); i++ {
            for j := 0; j < len(l2); j++ {
                if l2[j] != l1[i + j] {
                    break
                }
            	if j == len(l2) - 1 {
                    superlistBool = true
                }
            }
        }
    }
	switch {
    case sublistBool && superlistBool:
    	return equal
    case sublistBool && !superlistBool:
    	return sublist
    case !sublistBool && superlistBool:
    	return superlist
    default:
    	return unequal
    }
}
