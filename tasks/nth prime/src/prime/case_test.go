package prime

import (
	"testing"
)

var tests = []struct {
	description string
	input       int
	expected    int
	err         string
}{
	{
		description: "first prime",
		input:       1,
		expected:    2,
		err:         "",
	},
	{
		description: "second prime",
		input:       2,
		expected:    3,
		err:         "",
	},
	{
		description: "sixth prime",
		input:       6,
		expected:    13,
		err:         "",
	},
	{
		description: "big prime",
		input:       10001,
		expected:    104743,
		err:         "",
	},
	{
		description: "there is no zeroth prime",
		input:       0,
		expected:    0,
		err:         "there is no zeroth prime",
	},
}

func TestNth(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Nth(tc.input)

			switch {
			case tc.err != "":
				if err == nil {
					t.Fatalf("Nth(%d) expected error: %q, got: %d", tc.input, tc.err, actual)
				}

			case tc.expected != actual:
				t.Fatalf("Nth(%d) expected: %q, got: %q", tc.input, tc.expected, actual)

			case err != nil:
				t.Fatalf("Nth(%d) expected: %d, got error: %q", tc.input, tc.expected, err)
			}
		})
	}
}
