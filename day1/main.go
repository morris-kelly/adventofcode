package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read input.txt file
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	val := 50

	answer := 0

	// split by new line
	inputs := strings.SplitSeq(string(f), "\n")
	for input := range inputs {
		if input == "" {
			panic("blank string")
		}

		if strippedInput, ok := strings.CutPrefix(input, "L"); ok {
			val, answer = moveLeft(strippedInput, val, answer)
		} else if strippedInput, ok := strings.CutPrefix(input, "R"); ok {
			val, answer = moveRight(strippedInput, val, answer)
		} else {
			panic("invalid input")
		}
		fmt.Println("value is:", val)

		// part 1
		// if val == 0 {
		// 	answer++
		// }
	}

	println(answer)
}

func moveRight(strippedInput string, val, answer int) (int, int) {
	intVal := getValue(strippedInput)
	for range intVal {
		if val == 99 {
			answer++
			val = 0
		} else {
			val++
		}
	}
	return val, answer
}

func moveLeft(strippedInput string, val, answer int) (int, int) {
	intVal := getValue(strippedInput)
	for range intVal {
		switch val {
		case 1:
			// we hit zero, increment answer
			answer++
			val = 0
		case 0:
			val = 99
		default:
			val--
		}
	}
	return val, answer
}

func getValue(in string) int {
	intVal, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return intVal
}
