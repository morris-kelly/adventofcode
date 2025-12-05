package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type freshRange struct {
	min int
	max int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	freshRanges := []freshRange{}
	// freshIDs := []int{}

	lines := strings.SplitSeq(string(f), "\n")
	for line := range lines {

		if line == "" {
			// blank line, we've seen all the fresh ones
			break
		}

		switch strings.Contains(line, "-") {
		case true:
			// range of fresh IDs
			parts := strings.SplitN(line, "-", 2)
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}

			end, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			freshRanges = append(freshRanges, freshRange{min: start, max: end})

		case false:
			// 	// available ID
			// 	id, err := strconv.Atoi(line)
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	for _, val := range freshRanges {
			// 		if id >= val.min && id <= val.max {
			// 			freshIDs = append(freshIDs, id)
			// 			break
			// 		}
			// 	}
		}
	}

	filteredRanges := []freshRange{}

	for i, val := range freshRanges {
		switch i {
		case 0:
			// add the first one
			filteredRanges = append(filteredRanges, val)
		default:
			filteredRanges = mergeOverlappingRanges(filteredRanges, val)
		}
	}

	answer := 0
	for _, val := range filteredRanges {
		answer += (val.max - val.min) + 1
	}
	fmt.Println(answer)
}

func mergeOverlappingRanges(filteredRanges []freshRange, val freshRange) []freshRange {
	// Add the new range to the list
	allRanges := append(filteredRanges, val)

	// Keep merging until no more overlaps are found
	merged := true
	for merged {
		merged = false
		newRanges := []freshRange{}

		for i := 0; i < len(allRanges); i++ {
			current := allRanges[i]

			// Check if current range overlaps with any subsequent range
			for j := i + 1; j < len(allRanges); j++ {
				other := allRanges[j]

				// Check for overlap: ranges overlap if one starts before the other ends
				if current.min <= other.max && other.min <= current.max {
					// Merge the ranges
					current.min = min(current.min, other.min)
					current.max = max(current.max, other.max)

					// Remove the merged range from consideration
					allRanges = append(allRanges[:j], allRanges[j+1:]...)
					j-- // Adjust index after removal
					merged = true
				}
			}
			newRanges = append(newRanges, current)
		}
		allRanges = newRanges
	}

	// Sort ranges by min value
	for i := 0; i < len(allRanges); i++ {
		for j := i + 1; j < len(allRanges); j++ {
			if allRanges[i].min > allRanges[j].min {
				allRanges[i], allRanges[j] = allRanges[j], allRanges[i]
			}
		}
	}

	return allRanges
}
