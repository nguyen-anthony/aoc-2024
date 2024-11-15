package utils

func ContainsInt(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func UniqueInts(slice []int) []int {
	keys := make(map[int]bool)
	unique := []int{}
	for _, item := range slice {
		if _, found := keys[item]; !found {
			keys[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

