package main

import (
	"math"
	"testing"
)

var testCases = []struct {
	description string
	planet      Planet
	seconds     float64
	expected    float64
}{
	{
		description: "age on Earth",
		planet:      "Earth",
		seconds:     1000000000.0,
		expected:    31.69,
	},
	{
		description: "age on Mercury",
		planet:      "Mercury",
		seconds:     2134835688.0,
		expected:    280.88,
	},
	{
		description: "age on Venus",
		planet:      "Venus",
		seconds:     189839836.0,
		expected:    9.78,
	},
	{
		description: "age on Mars",
		planet:      "Mars",
		seconds:     2129871239.0,
		expected:    35.88,
	},
	{
		description: "age on Jupiter",
		planet:      "Jupiter",
		seconds:     901876382.0,
		expected:    2.41,
	},
	{
		description: "age on Saturn",
		planet:      "Saturn",
		seconds:     2000000000.0,
		expected:    2.15,
	},
	{
		description: "age on Uranus",
		planet:      "Uranus",
		seconds:     1210123456.0,
		expected:    0.46,
	},
	{
		description: "age on Neptune",
		planet:      "Neptune",
		seconds:     1821023456.0,
		expected:    0.35,
	},
	{
		description: "invalid planet causes error",
		planet:      "Sun",
		seconds:     680804807.0,
		expected:    -1.00,
	},
}

func TestAge(t *testing.T) {
	const precision = 0.01

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Age(tc.seconds, tc.planet)
			if math.Abs(actual-tc.expected) > precision {
				t.Fatalf("Age(%f, %v) = %f, want: %f", tc.seconds, tc.planet, actual, tc.expected)
			}
		})
	}
}

func BenchmarkAge(b *testing.B) {

	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Age(tc.seconds, tc.planet)
		}
	}
}
