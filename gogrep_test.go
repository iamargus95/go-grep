package gogrep

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	var nil1 []string
	var testCases = []struct {
		Files []string
		// Flag     []string
		Pattern  string
		Expected []string
	}{
		{
			Files:    []string{"This is One2N Consulting."},
			Pattern:  "Software",
			Expected: nil1,
		},
		{
			Files:    []string{"This is One2N Consulting."},
			Pattern:  "One2N",
			Expected: []string{"This is One2N Consulting."},
		},
		{
			Files:    []string{"This is One2N Consulting.", "Backend Engineers and Software Engineers are employed."},
			Pattern:  "Engineers",
			Expected: []string{"Backend Engineers and Software Engineers are employed."},
		},
		{
			Files:    []string{"This is One2N Consulting.", "One2N employs Backend Engineers and Software Engineers."},
			Pattern:  "One2N",
			Expected: []string{"This is One2N Consulting.", "One2N employs Backend Engineers and Software Engineers."},
		},
	}

	for _, tc := range testCases {
		actual := Grep(tc.Files, tc.Pattern)
		if !reflect.DeepEqual(actual, tc.Expected) {
			t.Errorf("Error Found.")
		}
	}
}
