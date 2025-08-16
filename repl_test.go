package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		}, {
			input:    " this is another test  ",
			expected: []string{"this", "is", "another", "test"},
		}, {
			input:    "Last Test  here  ",
			expected: []string{"last", "test", "here"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of slices do not match")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word does not match expected word")
			}
		}
	}
}
