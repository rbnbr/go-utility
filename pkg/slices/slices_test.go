package slices

import (
	"errors"
	"testing"
)

var (
	gotExpectedErrorFmt  = "got error: '%v' is not expected error: '%v'"
	gotExpectedResultFmt = "got result: '%v' is not equal expected result: '%v'"
)

// TestMakeUniqueStringSlice_success_1
// Tests the make unique string slice for a success case.
func TestMakeUniqueStringSlice_success_1(t *testing.T) {
	testSlice := []string{"hi", "hi", "Hodor", "Peter", "Pan", "Peter", "Rocco", "1", "1", "Peter", "0"}
	testSuffix := "_"

	expectedResult, expectedErr := []string{"hi", "hi_2", "Hodor", "Peter", "Pan", "Peter_2", "Rocco", "1", "1_2", "Peter_3", "0"}, ErrNil

	gotResult, gotError := MakeUniqueStringSlice(testSlice, testSuffix)
	if !errors.Is(gotError, expectedErr) {
		t.Errorf(gotExpectedErrorFmt, gotError, expectedErr)
	}

	if !Equal(expectedResult, gotResult, func(a string, b string) bool {
		return a == b
	}) {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestMakeUniqueStringSlice_failed_wrong_suffix_1
// Tests whether the method correctly throws the error on input which contains the suffix
func TestMakeUniqueStringSlice_failed_wrong_suffix_1(t *testing.T) {
	testSlice := []string{"hi", "hi", "Hodor", "hi_2"}
	testSuffix := "_"

	_, expectedErr := []string(nil), ErrInvalidArgument

	_, gotError := MakeUniqueStringSlice(testSlice, testSuffix)
	if !errors.Is(gotError, expectedErr) {
		t.Errorf(gotExpectedErrorFmt, gotError, expectedErr)
	}
}

// TestFindIndex_found
// Tests the find index function with integer slice and a predicate which should return something != -1, i.e., found the element
func TestFindIndex_found(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 5, 6}
	testPredicate := func(i int) bool {
		return testSlice[i] == 5
	}

	expectedResult := 4

	gotResult := FindIndex(len(testSlice), testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestFindIndex_not_found
// Tests the FindIndex function with a slice and predicate that should return -1
func TestFindIndex_not_found(t *testing.T) {
	testSlice := []string{"a", "b", "abc", "d"}
	testPredicate := func(i int) bool {
		return testSlice[i] == "god"
	}

	expectedResult := -1

	gotResult := FindIndex(len(testSlice), testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestFindIndexGeneric_found
// Tests the find index generic function with integer slice and a predicate which should return something != -1, i.e., found the element
func TestFindIndexGeneric_found(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 5, 6}
	testPredicate := func(i int) bool {
		return i == 6
	}

	expectedResult := 6

	gotResult := FindIndexGeneric(testSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestFindIndexGeneric_not_found
// Tests the FindIndexGeneric function with a slice and predicate that should return -1
func TestFindIndexGeneric_not_found(t *testing.T) {
	testSlice := []string{"a", "b", "abc", "d"}
	testPredicate := func(s string) bool {
		return s == "god"
	}

	expectedResult := -1

	gotResult := FindIndexGeneric(testSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestContainsGeneric
// Tests the ContainsGeneric function.
// One case true and one case false.
// Since this function directly wraps the Contains function, its test is omitted.
func TestContainsGeneric(t *testing.T) {
	testSlice := []int{1, 2, 3, 1, -1, 2, 1}

	testPredicate := func(i int) bool {
		return i == -1
	}

	expectedResult := true

	gotResult := ContainsGeneric(testSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}

	testPredicate = func(i int) bool {
		return i == 69
	}

	expectedResult = false

	gotResult = ContainsGeneric(testSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestCountGeneric
// Tests the CountGeneric function.
// One case it finds 3 elements, one case it finds none.
// Since this function directly wraps the Count function, its test is omitted.
func TestCountGeneric(t *testing.T) {
	testSlice := []int{1, 2, 3, -4, 1, -1, 2, 1}

	testPredicate := func(i int) bool {
		return i == 1
	}

	expectedResult := 3

	gotResult := CountGeneric(testSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}

	testPredicate = func(i int) bool {
		return i == -10
	}

	expectedResult = 0

	gotResult = CountGeneric(testSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestEqual
// Tests the Equal function.
// One case for equal and one case for not equal.
func TestEqual(t *testing.T) {
	testASlice := []int{1, 2, -1, 0, -4}
	testBSlice := []int{1, 2, -1, 0, -4}

	testCSlice := []int{4, 1, 0, 1, 5}
	testDSlice := []string(nil)
	testESlice := []int{}

	testPredicate := func(a, b int) bool {
		return a == b
	}

	// true cases

	expectedResult := true

	gotResult := Equal(testASlice, testBSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}

	testPredicate2 := func(a string, b int) bool {
		return false
	}

	gotResult = Equal(testDSlice, testESlice, testPredicate2)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}

	gotResult = Equal(testASlice, testASlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}

	// false cases
	expectedResult = false
	gotResult = Equal(testASlice, testCSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}

	gotResult = Equal(testDSlice, testASlice, testPredicate2)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}

	gotResult = Equal(testASlice, testESlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(gotExpectedResultFmt, gotResult, expectedResult)
	}
}
