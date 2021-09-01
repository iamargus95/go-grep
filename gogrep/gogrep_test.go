package gogrep

import (
	"reflect"
	"testing"
)

var nil1 []string

var testGrepNoFlag = []struct {
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

func TestGrep(t *testing.T) {

	for _, tc := range testGrepNoFlag {
		actual := Grep(tc.inputText, tc.pattern)
		if !reflect.DeepEqual(actual, tc.output) {
			t.Errorf("Error Found.")
		}
	}
}

var testGrepCaseSensitive = []struct {
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
		pattern:   "one2N",
		output:    []string{"This is One2N Consulting."},
	},
	{
		inputText: []string{"This is One2N Consulting.", "Backend Engineers and Software Engineers are employed."},
		pattern:   "engineers",
		output:    []string{"Backend Engineers and Software Engineers are employed."},
	},
	{
		inputText: []string{"This is One2N Consulting.", "One2N employs Backend Engineers and Software Engineers."},
		pattern:   "One2n",
		output:    []string{"This is One2N Consulting.", "One2N employs Backend Engineers and Software Engineers."},
	},
}

func TestGrepCaseSensitive(t *testing.T) {
	for _, tc := range testGrepCaseSensitive {
		actual := GrepCaseInsensitive(tc.inputText, tc.pattern)
		if !reflect.DeepEqual(actual, tc.output) {
			t.Errorf("Error Found.")
		}
	}
}

var testGrepCount = []struct {
	inputText []string
	pattern   string
	output    []string
}{
	{
		inputText: []string{"This is One2N Consulting."},
		pattern:   "Software",
		output:    []string{"0"},
	},
	{
		inputText: []string{"This is One2N Consulting."},
		pattern:   "One2N",
		output:    []string{"1"},
	},
	{
		inputText: []string{"This is One2N Consulting.", "Backend Engineers and Software Engineers are employed."},
		pattern:   "Engineers",
		output:    []string{"2"},
	},
	{
		inputText: []string{"This is One2N Consulting.", "One2N employs Backend Engineers and Software Engineers.", "I work at One2N"},
		pattern:   "One2N",
		output:    []string{"3"},
	},
}

func TestGrepCount(t *testing.T) {
	for _, tc := range testGrepCount {
		actual := GrepCount(tc.inputText, tc.pattern)
		if !reflect.DeepEqual(actual, tc.output) {
			t.Errorf("Error Found.")
		}
	}
}