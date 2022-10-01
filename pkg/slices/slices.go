package slices

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrNil             = error(nil)                     // nil error
	ErrInvalidArgument = errors.New("invalid argument") // there has been a problem with the provided argument
)

// MakeUniqueStringSlice
// Returns a copy of oldSlice where every element is unique.
// This is being done by going through the slice and counting the occurrence of the elements and setting each element which
// has occurred before to '{oldElement}{suffix}{count}', i.e., ["hello", "hello", "hello_"] with suffix "_"
// will return ["hello", "hello_2", "hello_"]
// If oldSlice contains elements that already match with regex: ^.*{suffix}[1-9]+$, it will return an error that the provided suffix cannot be used to make unique strings
func MakeUniqueStringSlice(oldSlice []string, suffix string) ([]string, error) {
	re := regexp.MustCompile(fmt.Sprintf("^.*%s[1-9]+$", suffix))
	if idx := FindIndexGeneric(oldSlice, func(s string) bool {
		return re.MatchString(s)
	}); idx != -1 {
		return nil, fmt.Errorf("%w: the provided suffix '%s' cannot be used to create unique strings because the provided slice contains elements which end with such suffix, e.g., '%s'", ErrInvalidArgument, suffix, oldSlice[idx])
	}

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
			ret[i] = fmt.Sprintf("%v%v%v", val, suffix, c+1)
		}
	}
	return ret, nil
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

// FindIndexGeneric
// Returns the index for which the provided predicate evaluates true first.
// Returns -1 if it didn't evaluate true at all.
// This is a wrapper around FindIndex using go generics.
// This omits the need to specify a limit or access the slice in the predicate.
// The limit is internally set to the full length of the slice and the now expected predicate directly uses the elements in the slice.
// Example:
// var someSlice []string // = ...
// idx := FindIndexGeneric(someSlice, func(s string) bool { return s == "Precious" }
func FindIndexGeneric[T any](slice []T, predicate func(T) bool) int {
	return FindIndex(len(slice), func(i int) bool {
		return predicate(slice[i])
	})
}

// Contains
// Returns true, if value is contained in the provided slice, else false.
// Similar usage as FindIndex
func Contains(limit int, predicate func(int) bool) bool {
	return FindIndex(limit, predicate) != -1
}

// ContainsGeneric
// Returns true, if value is contained in the provided slice, else false.
// Similar usage as FindIndexGeneric
func ContainsGeneric[T any](slice []T, predicate func(T) bool) bool {
	return FindIndex(len(slice), func(i int) bool {
		return predicate(slice[i])
	}) != -1
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

// CountGeneric
// Returns the amount of times the predicate did evaluate true
// Similar usage as FindIndexGeneric
func CountGeneric[T any](slice []T, predicate func(T) bool) int {
	return Count(len(slice), func(i int) bool {
		return predicate(slice[i])
	})
}

// Equal
// Given a predicate for comparison, returns true if the predicate evaluates true on all pairs of elements for the two slices,
// else false.
// Slices of same length evaluate to true, even though the one slice is nil while the other slice is an empty actual slice of possibly different type.
func Equal[T1 any, T2 any](t1 []T1, t2 []T2, predicate func(T1, T2) bool) bool {
	if len(t1) != len(t2) {
		return false
	}

	for i := 0; i < len(t1); i++ {
		if !predicate(t1[i], t2[i]) {
			return false
		}
	}

	return true
}

// TODO: implement remove duplicates
