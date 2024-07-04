package main

import "testing"

func TestClearInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths unequal: %v and %v", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			if word != c.expected[i] {
				t.Errorf("Words not equal: %v and  %v", word, c.expected[i])
			}
		}
	}

}
