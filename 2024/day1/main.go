package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// type set map[int]struct{}

func main() {
	file, err := os.Open("./in1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Part 1
	var leftList, rightList []int
	distance := 0
	// Part 2
	rightCountMap := map[int]int{}
	leftCountMap := map[int]int{}
	similarity := 0

	for scanner.Scan() {
		listNums := strings.Split(scanner.Text(), "   ")
		first, _ := strconv.Atoi(listNums[0])
		second, _ := strconv.Atoi(listNums[1])
		// Rather than append then sorting the list, keep a heap maybe?
		leftList = append(leftList, first)
		rightList = append(rightList, second)

		// For Part 2
		leftCountMap[first]++
		rightCountMap[second]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	for i := range len(leftList) {
		distance += intAbs(leftList[i] - rightList[i])
	}

	fmt.Println(distance)

	// Part 2
	for num, multiplier := range leftCountMap {
		if count, exists := rightCountMap[num]; exists {
			similarity += num * count * multiplier
		}
	}

	fmt.Println(similarity)
}

func intAbs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
