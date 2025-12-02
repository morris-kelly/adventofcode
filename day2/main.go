package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	day2()
}

func day2() {
	// read input.txt file
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	answer := 0

	ranges := strings.SplitSeq(string(f), ",")
	for val := range ranges {
		low, err := strconv.Atoi(strings.Split(val, "-")[0])
		if err != nil {
			panic(err)
		}
		high, err := strconv.Atoi(strings.Split(val, "-")[1])
		if err != nil {
			panic(err)
		}

		for i := low; i <= high; i++ {
			strI := strconv.Itoa(i)
			if hasRepeatedPattern(strI) {
				answer += i
			}
		}
	}
	println(answer)
}

func hasRepeatedPattern(s string) bool {
	patterns := make(map[string]bool)
	for i := 1; i <= len(s)/2; i++ {
		pattern := s[0:i]
		patterns[pattern] = true
	}

	for pattern := range patterns {
		repeated := true
		for j := 0; j < len(s); j += len(pattern) {
			if j+len(pattern) > len(s) || s[j:j+len(pattern)] != pattern {
				repeated = false
				break
			}
		}
		if repeated {
			return true
		}
	}

	return false
}
