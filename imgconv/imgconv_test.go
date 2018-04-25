package imgconv_test

import (
	"fmt"
	"testing"

	"github.com/execjosh/go-img-conv/imgconv"
)

var testTable = []struct {
	from        string
	to          string
	expectError bool
}{
	{"jpeg", "jpeg", true},
	{"foo", "bar", true},
	{"jpeg", "png", false},
	{"png", "jpeg", false},
}

func TestConversion(t *testing.T) {
	ic := imgconv.New("../testdata")
	for _, tt := range testTable {
		CheckResult(t, tt.from, tt.to, tt.expectError, ic.Convert(tt.from, tt.to))
	}
}

func CheckResult(t *testing.T, from, to string, expectError bool, err error) {
	t.Helper()

	if expectError && err == nil {
		t.Error(fmt.Sprint("Expected conversion from ", from, " to ", to, " to be rejected"))
	} else if !expectError && err != nil {
		t.Error(fmt.Sprint("Expected conversion from ", from, " to ", to, " to succeed"))
	}
}
