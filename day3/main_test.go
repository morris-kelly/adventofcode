package main

import "testing"

func TestGet12Digits(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"987654321111111", "987654321111"},
		{"811111111111119", "811111111119"},
		{"234234234234278", "434234234278"},
		{"818181911112111", "888911112111"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := getMaxDigits(test.input, 12)
			if result != test.expected {
				t.Errorf("For input %s, expected %s but got %s", test.input, test.expected, result)
			}
		})
	}
}
