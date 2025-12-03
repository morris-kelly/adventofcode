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
	data := string(f)

	answer := 0

	batteries := strings.SplitSeq(data, "\n")
	for battery := range batteries {
		biggest, secondBiggest := 0, 0
		for _, char := range battery {
			intVal, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			if intVal > biggest {
				secondBiggest = biggest
				biggest = intVal
			} else if intVal > secondBiggest {
				secondBiggest = intVal
			}
		}
		index := strings.Index(battery, strconv.Itoa(biggest))
		switch index {
		case 0:
			// the biggest is at the start, so the second biggest must be after it
			concat := strconv.Itoa(biggest) + strconv.Itoa(secondBiggest)
			val, err := strconv.Atoi(concat)
			if err != nil {
				panic(err)
			}
			answer += val
		case len(battery) - 1:
			// the biggest is at the end, so the second biggest must be before it
			concat := strconv.Itoa(secondBiggest) + strconv.Itoa(biggest)
			val, err := strconv.Atoi(concat)
			if err != nil {
				panic(err)
			}
			answer += val
		default:
			// the biggest is somewhere in the middle, so get the biggest after it
			biggestAfter := 0
			for _, char := range battery[index+1:] {
				intVal, err := strconv.Atoi(string(char))
				if err != nil {
					panic(err)
				}
				if intVal > biggestAfter {
					biggestAfter = intVal
				}
			}
			concat := strconv.Itoa(biggest) + strconv.Itoa(biggestAfter)
			val, err := strconv.Atoi(concat)
			if err != nil {
				panic(err)
			}
			answer += val
		}

	}
	fmt.Println("answer", answer)
}
