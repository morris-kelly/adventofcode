package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	freshRanges := []struct {
		min int
		max int
	}{}
	freshIDs := []int{}

	lines := strings.SplitSeq(string(f), "\n")
	for line := range lines {
		fmt.Println("line", line)

		if line == "" {
			// blank line, we've seen all the fresh ones
			continue
		}

		switch strings.Contains(line, "-") {
		case true:
			// range of fresh IDs
			parts := strings.SplitN(line, "-", 2)
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			fmt.Println("start", start)

			end, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			fmt.Println("end", end)

			freshRanges = append(freshRanges, struct {
				min int
				max int
			}{
				min: start,
				max: end,
			})
		case false:
			// available ID
			id, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			for _, val := range freshRanges {
				if id >= val.min && id <= val.max {
					freshIDs = append(freshIDs, id)
					break
				}
			}
		}
	}
	fmt.Println(len(freshIDs))
}
