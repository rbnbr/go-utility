package maps

import (
	"github.com/rbnbr/go-utility/pkg/consts"
	"github.com/rbnbr/go-utility/pkg/slices"
	"sort"
	"testing"
)

func TestGetKeysOfMap(t *testing.T) {
	testMap := map[string]int{"a": 0, "b": 2, "c": -3}

	expectedResult := []string{"a", "b", "c"}

	gotResult := GetKeysOfMap(testMap)

	// sort both since order may vary
	sort.Strings(gotResult)
	sort.Strings(expectedResult)

	if !slices.Equal(gotResult, expectedResult, func(t1 string, t2 string) bool {
		return t1 == t2
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}
}

func TestGetValuesOfMap(t *testing.T) {
	testMap := map[string]int{"a": 0, "b": 2, "c": -3}

	expectedResult := []int{0, 2, -3}

	gotResult := GetValuesOfMap(testMap)

	// sort both since order may vary
	sort.Ints(gotResult)
	sort.Ints(expectedResult)

	if !slices.Equal(gotResult, expectedResult, func(t1 int, t2 int) bool {
		return t1 == t2
	}) {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResult)
	}
}
