package function

import (
	"github.com/rbnbr/go-utility/pkg/consts"
	"testing"
)

func TestMakeGetReturnElementAt(t *testing.T) {
	testFunc := func() (int, int, uint, []rune, string) {
		return 1, 2, 3, nil, ""
	}

	expectedResultFirst := 1

	gotResult := GetFirstReturnElement(testFunc()).(int)

	if expectedResultFirst != gotResult {
		t.Errorf(consts.GotExpectedResultFmt, gotResult, expectedResultFirst)
	}

	expectedResultLast := ""

	gotResultLast := GetLastReturnElement(testFunc()).(string)
	if expectedResultLast != gotResultLast {
		t.Errorf(consts.GotExpectedResultFmt, gotResultLast, expectedResultLast)
	}
}
