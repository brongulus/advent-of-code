package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answer := 0
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		// safe := isSafe(levels) // Part 1
		safe := bruteDamp(levels) // Part 2
		if safe {
			answer++
		}
		// fmt.Printf("%v ", levels)
		// fmt.Println(safe)
	}

	fmt.Println(answer)
}

func isSafe(level []string) bool {
	isIncreasing := false
	isDecreasing := false
	for i := range len(level) - 1 {
		second, _ := strconv.Atoi(level[i+1])
		first, _ := strconv.Atoi(level[i])
		diff := second - first
		if diff < 0 {
			isDecreasing = true
			diff = -diff
		} else {
			isIncreasing = true
		}
		if isIncreasing == isDecreasing {
			return false
		}
		if !inRange(diff) {
			return false
		}
	}
	return true
} // 269

func bruteDamp(level []string) bool {
	if len(level) <= 2 || isSafe(level) {
		return true
	}
	for i := range len(level) {
		// https://stackoverflow.com/a/58726780 (WTF go)
		newLevel := append(level[:i:i], level[i+1:]...)
		if isSafe(newLevel) {
			return true
		}
	}
	return false
} // 337

func inRange(x int) bool {
	return x > 0 && x < 4
}

// func dampner(level []string) bool {
// 	if len(level) <= 2 {
// 		return true
// 	}
// 	signChange := 0
// 	damped := 0
// 	diffs := make([]int, len(level)-1)
// 	for i := 0; i < len(level)-2; i++ {
// 		third, _ := strconv.Atoi(level[i+2])
// 		second, _ := strconv.Atoi(level[i+1])
// 		first, _ := strconv.Atoi(level[i])
// 		diff := second - first
// 		diffOnDamp := third - first
// 		diffs[i] = diff
// 		fmt.Printf(" (%d: %d, %d, %d) ", i, diff, diffOnDamp, third-second)
// 		if diff*(third-second) <= 0 {
// 			signChange++
// 			i++
// 			// if (diffOnDamp == 0 || diff == 0 || (third-second) == 0) {
// 			// 	// check beg & end cases?
// 			// 	damped++
// 			// 	fmt.Println(diffs, signChange, "-ze-", damped)
// 			// 	return false
// 			// }
// 			// continue
// 		}
// 		diff = max(diff, -diff)
// 		diffOnDamp = max(diffOnDamp, -diffOnDamp)
// 		if !inRange(diff) {
// 			if inRange(diffOnDamp) {
// 				damped++
// 				i++
// 			} else {
// 				fmt.Println(diffs, signChange, "-ee-", damped)
// 				return false
// 			}
// 		}
// 		if signChange + damped > 1 {
// 			fmt.Println(diffs, signChange, damped)
// 			return false
// 		}
// 	}

// 	fmt.Println(diffs, signChange, damped)
// 	return true
// }
