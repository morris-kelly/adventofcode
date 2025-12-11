package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	data := string(f)

	length := 0
	signs := []string{}
	answers := []int{}

	grid := [][]string{}

	split := strings.Split(data, "\n")
	slices.Reverse(split)
	for i, v := range split {

		nums := strings.Split(v, " ")
		switch i {
		case 0:
			// the signs are on line 1 now
			for _, v := range nums {
				if v != "" {
					signs = append(signs, v)

					// initialise answers
					answers = append(answers, 0)
				}
			}

			// store length for later
			length = len(signs)

		default:
			temp := []string{}
			maxLength := 0

			// we now have numbers, so use the sign and do the maths
			for _, v := range nums {
				if v != "" {
					temp = append(temp, v)
					if len(v) > maxLength {
						maxLength = len(v)
					}
				}
			}
			if len(temp) != length {
				panic("lengths don't match")
			}
			grid = append(grid, temp)

		}

	}

	// now we have the grid, we can process it
	// loop through each column and find the longest number
	for col := 0; col < length; col++ {
		maxLength := 0
		for row := range grid {
			if len(grid[row][col]) > maxLength {
				maxLength = len(grid[row][col])
			}
		}
		fmt.Println("max length", maxLength)

		temp := []string{}
		for row := 0; row < len(grid); row++ {
			number := ""
			// use the max length, count down and build the number
			for i := maxLength - 1; i >= 0; i-- {
				if len(grid[row][col]) <= i {
					continue
				}
				number += string(grid[row][col][i])
				fmt.Println("number", number)
			}
			temp = append(temp, number)
		}
		fmt.Println("temp", temp)
	}

	// finally, sum the answers
	finalAnswer := 0
	for _, answer := range answers {
		finalAnswer += answer
	}
	fmt.Println("final answer", finalAnswer)
}
