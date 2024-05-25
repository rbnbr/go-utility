package slices

import (
	"errors"
	"github.com/rbnbr/go-utility/pkg/consts"
	"github.com/rbnbr/go-utility/pkg/maps"
	"sort"
	"strings"
	"testing"
)

// TestMakeUniqueStringSlice_success_1
// Tests the make unique string slice for a success case.
func TestMakeUniqueStringSlice_success_1(t *testing.T) {
	testSlice := []string{"hi", "hi", "Hodor", "Peter", "Pan", "Peter", "Rocco", "1", "1", "Peter", "0"}
	testSuffix := "_"

	expectedResult, expectedErr := []string{"hi", "hi_2", "Hodor", "Peter", "Pan", "Peter_2", "Rocco", "1", "1_2", "Peter_3", "0"}, ErrNil

	gotResult, gotError := MakeUniqueStringSlice(testSlice, testSuffix)
	if !errors.Is(gotError, expectedErr) {
		t.Errorf(consts.GotExpectedErrorFmt, gotError, expectedErr)
	}

	if !Equal(expectedResult, gotResult, func(a string, b string) bool {
		return a == b
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
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
		t.Errorf(consts.GotExpectedErrorFmt, gotError, expectedErr)
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
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
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
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
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
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
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
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
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
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testPredicate = func(i int) bool {
		return i == 69
	}

	expectedResult = false

	gotResult = ContainsGeneric(testSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
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
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testPredicate = func(i int) bool {
		return i == -10
	}

	expectedResult = 0

	gotResult = CountGeneric(testSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
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
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testPredicate2 := func(a string, b int) bool {
		return false
	}

	gotResult = Equal(testDSlice, testESlice, testPredicate2)
	if gotResult != expectedResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	gotResult = Equal(testASlice, testASlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	// false cases
	expectedResult = false
	gotResult = Equal(testASlice, testCSlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	gotResult = Equal(testDSlice, testASlice, testPredicate2)
	if gotResult != expectedResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	gotResult = Equal(testASlice, testESlice, testPredicate)
	if gotResult != expectedResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestUniqueConfigurable
// Tests UniqueConfigurable and it's variations on float slices.
func TestUniqueConfigurable(t *testing.T) {
	testSlice := []float64{1.0, 2.0, 4.0, -1.0, -2.0, 5.0, 1.1, 2.1, 0.0, 1.2}
	testPredicate := func(a, b float64) bool {
		return int(a) == int(b) // rounding down to be able to differentiate between 'same' elements
	}

	// unique (first occurrence)
	expectedResult := []float64{1.0, 2.0, 4.0, -1.0, -2.0, 5.0, 0.0}
	gotResult := Unique(testSlice, testPredicate)

	if !Equal(gotResult, expectedResult, func(a, b float64) bool {
		return a == b
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	// unique (last occurrence)
	expectedResult = []float64{1.2, 2.1, 4.0, -1.0, -2.0, 5.0, 0.0}
	gotResult = UniqueLast(testSlice, testPredicate)

	if !Equal(gotResult, expectedResult, func(a, b float64) bool {
		return a == b
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}
}

// TestFilter
// Tests the generic Filter function.
func TestFilter(t *testing.T) {
	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testPredicate := func(i int) bool { return i%2 == 0 }

	expectedResult := []int{0, 2, 4, 6, 8}
	gotResult := Filter(testSlice, testPredicate)

	if !Equal(gotResult, expectedResult, func(a, b int) bool {
		return a == b
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testSliceString := []string{"hallo", "hello", "alo", "vera"}
	testPredicateString := func(s string) bool {
		return s[0] == 'a' || s[0] == 'v'
	}

	expectedResultString := []string{"alo", "vera"}
	gotResultString := Filter(testSliceString, testPredicateString)

	if !Equal(gotResultString, expectedResultString, strings.EqualFold) {
		t.Errorf(consts.GotExpectedResultFmt, gotResultString, expectedResultString)
	}
}

// TestMap
// Tests the generic Map function.
func TestMap(t *testing.T) {
	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testMapping := func(i *int) int {
		return (*i) * 2
	}

	expectedResult := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	gotResult := Map(testSlice, testMapping)

	if !Equal(gotResult, expectedResult, func(a int, b int) bool {
		return a == b
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testSliceString := []string{"hallo", "hello", "alo", "vera"}
	testMappingString := func(s *string) int {
		// count number of letters
		return len(*s)
	}

	expectedResultString := []int{5, 5, 3, 4}
	gotResultString := Map(testSliceString, testMappingString)

	if !Equal(gotResultString, expectedResultString, func(a int, b int) bool {
		return a == b
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}
}


// TestMapI
// Tests the generic MapI function.
func TestMapI(t *testing.T) {
	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testMapping := func(i *int, idx int) int {
		return (*i) * 2 + idx
	}

	expectedResult := []int{0, 3, 6, 9, 12, 15, 18, 21, 24, 27}
	gotResult := MapI(testSlice, testMapping)

	if !Equal(gotResult, expectedResult, func(a int, b int) bool {
		return a == b
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}
}

func TestReduce(t *testing.T) {
	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testReduceFunc := func(i *int, v *int) int {
		return (*i) + (*v)
	}

	expectedResult := 45
	gotResult := Reduce(testSlice, testReduceFunc, 0)

	if gotResult != expectedResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testSliceConcat := [][]int{{0, 1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}}
	testReduceFuncConcat := func(i *[]int, v *[]int) []int {
		// concat both
		return append(*v, *i...)
	}

	expectedResultConcat := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	gotResultConcat := Reduce(testSliceConcat, testReduceFuncConcat, []int{})

	if !Equal(gotResultConcat, expectedResultConcat, func(i int, i2 int) bool {
		return i == i2
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResultConcat, expectedResultConcat)
	}
}

func TestConcatSlices(t *testing.T) {
	testSliceConcat := [][]int{{0, 1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}}

	expectedResultConcat := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	gotResultConcat := ConcatSlices(testSliceConcat)

	if !Equal(gotResultConcat, expectedResultConcat, func(i int, i2 int) bool {
		return i == i2
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResultConcat, expectedResultConcat)
	}
}

func TestAny(t *testing.T) {
	testSlice := []bool{false, false, false, false}

	expectedResult := false

	gotResult := Any(testSlice)

	if expectedResult != gotResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testSlice2 := []bool{false, true, false, false}

	expectedResult2 := true

	gotResult2 := Any(testSlice2)

	if expectedResult2 != gotResult2 {
		t.Errorf(consts.GotExpectedResultFmt, gotResult2, expectedResult2)
	}
}

func TestAll(t *testing.T) {
	testSlice := []bool{true, true, false, true}

	expectedResult := false

	gotResult := All(testSlice)

	if expectedResult != gotResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testSlice2 := []bool{true, true, true, true}

	expectedResult2 := true

	gotResult2 := All(testSlice2)

	if expectedResult2 != gotResult2 {
		t.Errorf(consts.GotExpectedResultFmt, gotResult2, expectedResult2)
	}
}

func TestGroupBy(t *testing.T) {
	testSliceToGroup := []int{0, 1, 2, 3, 1, 2, 30, 1, -1, -2, -1, 3, 2, 7, -3, 19, 39, 10, 20}

	expectedResult := map[int][]int{
		0:  {0},
		1:  {1, 1, 1},
		2:  {2, 2, 2},
		3:  {3, 3},
		30: {30},
		-1: {-1, -1},
		-2: {-2},
		7:  {7},
		-3: {-3},
		19: {19},
		39: {39},
		10: {10},
		20: {20},
	}

	gotResult := GroupBy(testSliceToGroup, func(v int) int {
		return v
	})

	gotResultKeys := maps.GetKeysOfMap(gotResult)
	expectedResultKeys := maps.GetKeysOfMap(expectedResult)

	sort.Ints(gotResultKeys)
	sort.Ints(expectedResultKeys)

	if !Equal(gotResultKeys, expectedResultKeys, func(i int, i2 int) bool {
		return i == i2 && Equal(gotResult[i], expectedResult[i], func(i3 int, i4 int) bool {
			return i3 == i4
		})
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}

	testSliceToGroup2 := []int{0, 1, 2, 3, 1, 2, 30, 1, -1, -2, -1, 3, 2, 7, -3, 19, 39, 10, 20}

	expectedResult2 := map[int][]int{
		0:  {0},
		1:  {1, 1, 1},
		2:  {2, 2, 2},
		3:  {3, 3},
		30: {30},
		-1: {-1, -1},
		-2: {-2},
		7:  {7},
		-3: {-3},
		19: {19},
		39: {39},
		10: {10},
		20: {20},
		34: {34},
	}

	gotResult2 := GroupBy(testSliceToGroup2, func(v int) int {
		return v
	})

	gotResultKeys2 := maps.GetKeysOfMap(gotResult2)
	expectedResultKeys2 := maps.GetKeysOfMap(expectedResult2)

	sort.Ints(gotResultKeys2)
	sort.Ints(expectedResultKeys2)

	if Equal(gotResultKeys2, expectedResultKeys2, func(i int, i2 int) bool {
		return i == i2 && Equal(gotResult2[i], expectedResult2[i], func(i3 int, i4 int) bool {
			return i3 == i4
		})
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult2, expectedResult2)
	}
}
