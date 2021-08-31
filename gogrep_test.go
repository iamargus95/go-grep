package gogrep

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	var nil1 []string
	var testCases = []struct {
		Files    []byte
		Pattern  string
		Expected []string
	}{
		{
			Files:    []byte("This is One2N Consulting."),
			Pattern:  "Software",
			Expected: nil1,
		},
		{
			Files:    []byte("This is One2N Consulting."),
			Pattern:  "One2N",
			Expected: []string{"This is One2N Consulting.\n"},
		},
		{
			Files:    []byte("This is One2N Consulting. We employ Backend Engineers and Software Engineers."),
			Pattern:  "Engineers",
			Expected: []string{" We employ Backend Engineers and Software Engineers.\n"},
		},
		// {
		// 	Files:    []byte("This is One2N Consulting.One2N employs Backend Engineers and Software Engineers."),
		// 	Pattern:  "One2N",
		// 	Expected: []string{"This is One2N Consulting.\n One2N employs Backend Engineers and Software Engineers.\n"},
		// }, -------------------> Need for formatting arises.
	}

	for _, tc := range testCases {
		actual := Grep(tc.Files, tc.Pattern)
		if !reflect.DeepEqual(actual, tc.Expected) {
			t.Errorf("Error Found.")
		}
	}
}
