package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answer := 0
	// Ref: https://pkg.go.dev/regexp/syntax
	// mul, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`) // Part 1
	updatedMul, _ := regexp.Compile(`(mul\([0-9]+,[0-9]+\)|do(n't)??\(\))`) // Part 2
	enabled := true                                                         // Part 2
	findNum, _ := regexp.Compile(`([0-9]+)`)

	for scanner.Scan() {
		text := scanner.Text()
		// matches := mul.FindAllString(text, -1) // Part 1
		matches := updatedMul.FindAllString(text, -1) // Part 2

		for _, match := range matches {
			if match == "do()" {
				enabled = true
				continue
			} else if match == "don't()" {
				enabled = false
				continue
			}
			if enabled {
				nums := findNum.FindAllString(match, -1)
				first, _ := strconv.Atoi(nums[0])
				second, _ := strconv.Atoi(nums[1])
				answer += first * second
			}
		}
	}

	fmt.Println(answer)
}
