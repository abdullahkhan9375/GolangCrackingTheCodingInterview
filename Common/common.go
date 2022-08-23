package common

func IsEqual(aArray1 []int, aArray2 []int) bool {
	var lSecondIndex = 0
	for _, aVal := range aArray1 {
		if aVal != aArray2[lSecondIndex] {
			return false
		}
		lSecondIndex++
	}
	return true
}
