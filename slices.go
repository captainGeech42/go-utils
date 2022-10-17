package utils

// Check if a slice contains a specified element
func Contains[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

// Check if a slice contains a specified element, only between a certain range
func ContainsAfter[T comparable](haystack []T, needle T, afterIdx int) bool {
	for _, v := range haystack[afterIdx:] {
		if v == needle {
			return true
		}
	}

	return false
}

// Check if a slice contains a specified element, only between a certain range
func ContainsWithinRange[T comparable](haystack []T, needle T, lowerBound int, upperBound int) bool {
	for _, v := range haystack[lowerBound:upperBound] {
		if v == needle {
			return true
		}
	}

	return false
}
