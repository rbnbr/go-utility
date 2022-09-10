package slices

import "fmt"

// MakeUniqueStringSlice
// Returns a copy of oldSlice where every element is unique.
// This is being done by going through the slice and counting the occurrence of the elements and setting each element which
// has occurred before to '{oldElement}{postfix}{count}', i.e., ["hello", "hello", "hello_"] with postfix "_"
// will return ["hello", "hello_2", "hello_"]
func MakeUniqueStringSlice(oldSlice []string, postfix string) []string {
	ret := make([]string, len(oldSlice))
	counts := make(map[string]int)
	for i, val := range oldSlice {
		c, ok := counts[val]
		if !ok {
			// first time
			counts[val] = 1
			ret[i] = val
		} else {
			// seen before
			counts[val] = c + 1
			ret[i] = fmt.Sprintf("%v%v%v", val, postfix, c+1)
		}
	}
	return ret
}

// FindIndex
// Returns the index for which the provided predicate evaluates true first.
// Returns -1 if it didn't evaluate true at all.
// Limit specifies the search space for the indices which is [0, limit-1]
// No bounds check included.
// Example:
// var someSlice []string // = ...
// idx := FindIndex(len(someSlice), func(i int) bool { return someSlice[i] == "Precious" }
func FindIndex(limit int, predicate func(int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// Contains
// Returns true, if value is contained in the provided slice, else false.
// Similar usage as FindIndex
func Contains(limit int, predicate func(int) bool) bool {
	return FindIndex(limit, predicate) != -1
}

// Count
// Returns the amount of times the predicate did evaluate true
// Similar usage as FindIndex
func Count(limit int, predicate func(int) bool) int {
	count := 0
	for i := 0; i < limit; i++ {
		if predicate(i) {
			count += 1
		}
	}
	return count
}
