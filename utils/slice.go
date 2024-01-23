package utils

func ElementInSlice(element string, slice []string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func RemoveElementFromSlice(element string, slice []string) []string {
	for i, e := range slice {
		if e == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
