package cryptosquare

import "testing"

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "Test for special characters",
		input:       "s#$%^&plunk",
		expected:    "su pn lk",
	},
	{
		description: "Test for special characters",
		input:       "Time is an illusion. Lunchtime doubly so.",
		expected:    "tasney inicds miohoo elntu  illib  suuml ",
	},
	{
		input:    "s#$%^&plunk",
		expected: "su pn lk",
	},
	{
		input:    "1, 2, 3 GO!",
		expected: "1g 2o 3 ",
	},

	{
		input:    "1234",
		expected: "13 24",
	},
	{
		input:    "123456789",
		expected: "147 258 369",
	},
	{
		input:    "123456789abc",
		expected: "159 26a 37b 48c",
	},
	{
		input:    "Never vex thine heart with idle woes",
		expected: "neewl exhie vtetw ehaho ririe vntds",
	},
	{
		input:    "ZOMG! ZOMBIES!!!",
		expected: "zzi ooe mms gb ",
	},
	{
		input:    "Time is an illusion. Lunchtime doubly so.",
		expected: "tasney inicds miohoo elntu  illib  suuml ",
	},
	{
		input:    "We all know interspecies romance is weird.",
		expected: "wneiaw eorene awssci liprer lneoid ktcms ",
	},
	{
		input:    "12345678",
		expected: "147 258 36 ",
	},

	{
		input:    "123456789a",
		expected: "159 26a 37  48 ",
	},
	{
		input:    "If man was meant to stay on the ground god would have given us roots",
		expected: "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau ",
	},
	{
		input:    "Have a nice day. Feed the dog & chill out!",
		expected: "hifei acedl veeol eddgo aatcu nyhht",
	},
}

func TestForSpecialCharacters(t *testing.T) {
	for _, test_case := range testCases {
		t.Run(test_case.description, func(t *testing.T) {
			actual := Encode(test_case.input)
			if actual != test_case.expected {
				t.Fatalf("Expected \"%v\", got \"%v\"", test_case.expected, actual)
			}
		})
	}
}
