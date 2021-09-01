package gogrep

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	var nil1 []string
	var testCases = []struct {
		inputText []string
		pattern   string
		output    []string
	}{
		{
			inputText: []string{"This is One2N Consulting."},
			pattern:   "Software",
			output:    nil1,
		},
		{
			inputText: []string{"This is One2N Consulting."},
			pattern:   "One2N",
			output:    []string{"This is One2N Consulting."},
		},
		{
			inputText: []string{"This is One2N Consulting.", "Backend Engineers and Software Engineers are employed."},
			pattern:   "Engineers",
			output:    []string{"Backend Engineers and Software Engineers are employed."},
		},
		{
			inputText: []string{"This is One2N Consulting.", "One2N employs Backend Engineers and Software Engineers."},
			pattern:   "One2N",
			output:    []string{"This is One2N Consulting.", "One2N employs Backend Engineers and Software Engineers."},
		},
	}

	for _, tc := range testCases {
		actual := Grep(tc.inputText, tc.pattern)
		if !reflect.DeepEqual(actual, tc.output) {
			t.Errorf("Error Found.")
		}
	}
}
