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
// Two slices of length 0 evaluate to true, even though the one slice is nil while the other slice is an empty actual slice of possibly different type.
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

// UniqueConfigurable
// Returns a new slice with only elements which do not occur twice using the predicate function to compare equality between two elements.
// Is configurable to return the first or the last occurrence per element.
// When chosen last occurrence, then the order is still the same as with first occurrence but the elements has been replaced.
// Has quadratic runtime due to usage of slice for lookup and not map which allows elements that are non-comparable
func UniqueConfigurable[T any](slice []T, predicate func(T, T) bool, firstOccurrence bool) []T {
	unique := make([]T, 0)

	for i := 0; i < len(slice); i++ {
		if idx := FindIndexGeneric(unique, func(u T) bool {
			return predicate(slice[i], u)
		}); idx == -1 {
			// current element in slice, i.e., slice[i], does not exist in unique slice. -> can be added
			unique = append(unique, slice[i])
		} else {
			// current element in slice, i.e., slice[i], does exist in unique slice
			if firstOccurrence {
				// we only keep the first occurrence -> do nothing
				continue
			} else {
				// we keep the last occurrence -> replace element in unique with current element
				unique[idx] = slice[i]
			}
		}
	}

	return unique
}

// Unique
// Shorthand for UniqueFirst(..)
func Unique[T any](slice []T, predicate func(T, T) bool) []T {
	return UniqueFirst(slice, predicate)
}

// UniqueFirst
// Shorthand for UniqueConfigurable(slice, predicate, true)
func UniqueFirst[T any](slice []T, predicate func(T, T) bool) []T {
	return UniqueConfigurable(slice, predicate, true)
}

// UniqueLast
// Shorthand for UniqueConfigurable(slice, predicate, false)
func UniqueLast[T any](slice []T, predicate func(T, T) bool) []T {
	return UniqueConfigurable(slice, predicate, false)
}

// Filter
// Returns a new slice t which only contains those elements of the provided slice
// for which the given predicate function evaluates to true.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	ret := make([]T, 0)

	for _, v := range slice {
		if predicate(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

// Map
// Maps a slice to another slice by applying a function 'mapping' to each of its elements.
func Map[T any, V any](slice []T, mapping func(t *T) V) []V {
	ret := make([]V, len(slice))

	for i := 0; i < len(slice); i++ {
		ret[i] = mapping(&slice[i])
	}

	return ret
}
