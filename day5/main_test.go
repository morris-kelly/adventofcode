package main

import "testing"

func TestMergeOverlappingRanges(t *testing.T) {
	tests := []struct {
		name     string
		ranges   []freshRange
		newRange freshRange
		expected []freshRange
	}{
		{
			name:     "no overlap",
			ranges:   []freshRange{{min: 1, max: 5}, {min: 10, max: 15}},
			newRange: freshRange{min: 6, max: 9},
			expected: []freshRange{{min: 1, max: 5}, {min: 6, max: 9}, {min: 10, max: 15}},
		},
		{
			name:     "overlap with one range",
			ranges:   []freshRange{{min: 1, max: 5}, {min: 10, max: 15}},
			newRange: freshRange{min: 4, max: 12},
			expected: []freshRange{{min: 1, max: 15}},
		},
		{
			name:     "overlap with multiple ranges",
			ranges:   []freshRange{{min: 1, max: 5}, {min: 10, max: 15}, {min: 20, max: 25}},
			newRange: freshRange{min: 4, max: 22},
			expected: []freshRange{{min: 1, max: 25}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeOverlappingRanges(tt.ranges, tt.newRange)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %v, got %v", tt.expected, result)
					return
				}
			}
		})
	}
}
