package sublist

type Relation string

func Sublist(list1, list2 []int) Relation {
	if len(list1) == len(list2) && isSub(list1, list2) {
		return "equal"
	}
	if len(list1) > len(list2) && isSub(list2, list1) {
		return "superlist"
	}
	if len(list2) > len(list1) && isSub(list1, list2) {
		return "sublist"
	}
	return "unequal"
}

func isSub(sm, la []int) bool {
	for i := 0; i <= len(la)-len(sm); i++ {
		match := true
		for j := 0; j < len(sm); j++ {
			if sm[j] != la[i+j] {
				match = false
				break
			}
		}

		if match {
			return true
		}

	}
	return false
}
