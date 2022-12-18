package slice

func Remove[T comparable](s []T, elem T) []T {
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
