package phonenumber

import (
	"testing"
)

type testCase struct {
	description       string
	input             string
	expectErr         bool
	expectedNumber    string
	expectedAreaCode  string
	expectedFormatted string
}

var testCases = []testCase{
	{
		description:       "cleans the number",
		input:             "(223) 456-7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "cleans numbers with dots",
		input:             "223.456.7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "cleans numbers with multiple spaces",
		input:             "223 456   7890   ",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "invalid when 9 digits",
		input:             "123456789",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid when 11 digits does not start with a 1",
		input:             "22234567890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "valid when 11 digits and starting with 1",
		input:             "12234567890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "valid when 11 digits and starting with 1 even with punctuation",
		input:             "+1 (223) 456-7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "invalid when more than 11 digits",
		input:             "321234567890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid with letters",
		input:             "523-abc-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid with punctuations",
		input:             "523-@:!-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 0",
		input:             "(023) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 1",
		input:             "(123) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 0",
		input:             "(223) 056-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 1",
		input:             "(223) 156-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 0 on valid 11-digit number",
		input:             "1 (023) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 1 on valid 11-digit number",
		input:             "1 (123) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 0 on valid 11-digit number",
		input:             "1 (223) 056-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 1 on valid 11-digit number",
		input:             "1 (223) 156-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
}

func TestNumber(t *testing.T) {
	runTests("Number", Number, func(tc testCase) string { return tc.expectedNumber }, t)
}

func runTests(
	funcName string,
	f func(s string) (string, error),
	getExpected func(tc testCase) string,
	t *testing.T,
) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, actualErr := f(tc.input)
			switch {
			case tc.expectErr:
				if actualErr == nil {
					t.Fatalf("%s(%q) expected error, got: %q", funcName, tc.input, actual)
				}

			case actualErr != nil:
				t.Fatalf("%s(%q) returned error: %v, want: %q", funcName, tc.input, actualErr, getExpected(tc))
			case actual != getExpected(tc):
				t.Fatalf("%s(%q) = %q, want: %q", funcName, tc.input, actual, getExpected(tc))
			}
		})
	}
}

func BenchmarkFormat(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Format(test.input)
		}
	}
}

func TestFormat(t *testing.T) {
	runTests("Format", Format, func(tc testCase) string { return tc.expectedFormatted }, t)
}

func BenchmarkNumber(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Number(test.input)
		}
	}
}

func TestAreaCode(t *testing.T) {
	runTests("AreaCode", AreaCode, func(tc testCase) string { return tc.expectedAreaCode }, t)
}

func BenchmarkAreaCode(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			AreaCode(test.input)
		}
	}
}
