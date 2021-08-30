package gogrep

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Files    []byte
	Pattern  string
	Expected []string
}{
	{
		Files:    []byte("This is One2N Consulting."),
		Pattern:  "One2N",
		Expected: []string{"One2N"},
	},
	// {
	// 	Files:    []byte("This is One2N Consulting."),
	// 	Pattern:  "Software",
	// 	Expected: []string{""},
	// },
	{
		Files:    []byte("This is One2N Consulting. We employ Backend Engineers and Software Engineers"),
		Pattern:  "Engineers",
		Expected: []string{"Engineers Engineers"},
	},
}

func TestGrep(t *testing.T) {
	for _, tc := range testCases {
		actual := Grep(tc.Files, tc.Pattern)
		if !reflect.DeepEqual(actual, tc.Expected) {
			t.Errorf("Error Found.")
		}
	}
}
