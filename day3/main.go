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
		// get the biggest 12 digits from the battery
		twelveDigits := getMaxDigits(battery, 12)

		val, err := strconv.Atoi(twelveDigits)
		if err != nil {
			panic(err)
		}
		answer += val

		// index := strings.Index(battery, strconv.Itoa(biggest))
		// switch index {
		// case 0:
		// 	// the biggest is at the start, so the second biggest must be after it
		// 	concat := strconv.Itoa(biggest) + strconv.Itoa(secondBiggest)
		// 	val, err := strconv.Atoi(concat)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	answer += val
		// case len(battery) - 1:
		// 	// the biggest is at the end, so the second biggest must be before it
		// 	concat := strconv.Itoa(secondBiggest) + strconv.Itoa(biggest)
		// 	val, err := strconv.Atoi(concat)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	answer += val
		// default:
		// 	// the biggest is somewhere in the middle, so get the biggest after it
		// 	biggestAfter := 0
		// 	for _, char := range battery[index+1:] {
		// 		intVal, err := strconv.Atoi(string(char))
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		if intVal > biggestAfter {
		// 			biggestAfter = intVal
		// 		}
		// 	}
		// 	concat := strconv.Itoa(biggest) + strconv.Itoa(biggestAfter)
		// 	val, err := strconv.Atoi(concat)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	answer += val
		// }

	}
	fmt.Println("answer", answer)
}

func getMaxDigits(input string, digits int) string {
	ints := []int{}
	for _, char := range input {
		intVal, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		ints = append(ints, intVal)
	}

	// iterate through the ints and build the biggest number with the given digits
	result := []int{}
	toRemove := len(ints) - digits
	for i := 0; i < len(ints); i++ {
		// while we can still remove digits and the last digit in result is less than the current digit
		for toRemove > 0 && len(result) > 0 && result[len(result)-1] < ints[i] {
			// remove the last digit from result
			result = result[:len(result)-1]
			toRemove--
		}
		result = append(result, ints[i])
	}

	// if we still have to remove digits, remove them from the end
	result = result[:len(result)-toRemove]

	concat := ""
	for _, digit := range result {
		concat += strconv.Itoa(digit)
	}
	return concat
}
