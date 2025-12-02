package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveLeft(t *testing.T) {
	tests := []struct {
		strippedInput string
		val           int
		want          int
	}{
		{"10", 50, 0},
		{"60", 50, 1},
		{"51", 50, 1},
		{"50", 50, 1},
		{"49", 50, 0},
	}

	for _, tt := range tests {
		t.Run(tt.strippedInput, func(t *testing.T) {
			_, answer := moveLeft(tt.strippedInput, tt.val, 0)
			assert.Equal(t, tt.want, answer)
		})
	}
}

func TestMoveRight(t *testing.T) {
	tests := []struct {
		strippedInput string
		val           int
		want          int
	}{
		{"10", 50, 0},
		{"49", 50, 0},
		{"50", 50, 1},
		{"51", 50, 1},
		{"60", 50, 1},
		{"1000", 50, 10},
	}

	for _, tt := range tests {
		t.Run(tt.strippedInput, func(t *testing.T) {
			_, answer := moveRight(tt.strippedInput, tt.val, 0)
			assert.Equal(t, tt.want, answer)
		})
	}
}
