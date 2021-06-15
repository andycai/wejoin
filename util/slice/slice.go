package slice

func RemoveInt(s []int, elem int) []int {
	index := -1
	for i, v := range s {
		if v == elem {
			index = i
			break
		}
	}
	newS := append(s[:index], s[index:]...)
	return newS
}

func RemoveI64(s []int64, elem int64) []int64 {
	index := -1
	for i, v := range s {
		if v == elem {
			index = i
			break
		}
	}
	newS := append(s[:index], s[index:]...)
	return newS
}

func RemoveU64(s []uint64, elem uint64) []uint64 {
	index := -1
	for i, v := range s {
		if v == elem {
			index = i
			break
		}
	}
	newS := append(s[:index], s[index:]...)
	return newS
}
