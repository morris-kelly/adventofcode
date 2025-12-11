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

	split := strings.Split(data, "\n")
	slices.Reverse(split)
	for i, v := range split {
		nums := strings.Split(v, " ")
		fmt.Println("nums ", nums)
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
			length = len(signs)
		default:
			temp := []string{}
			// we now have numbers, so use the sign and do the maths
			for _, v := range nums {
				if v != "" {
					temp = append(temp, v)
				}
			}
			if len(temp) != length {
				panic("lengths don't match")
			}

			for i := range length {
				switch signs[i] {
				case "+":
					num := 0
					fmt.Sscanf(temp[i], "%d", &num)
					answers[i] += num
				case "*":
					num := 0
					fmt.Sscanf(temp[i], "%d", &num)
					if answers[i] == 0 {
						answers[i] = 1
					}
					answers[i] *= num
				default:
					panic("unknown sign " + signs[i])
				}
			}
		}

	}
	finalAnswer := 0
	for _, answer := range answers {
		finalAnswer += answer
	}
	fmt.Println("final answer", finalAnswer)
}
