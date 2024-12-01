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
// Src: gopkg
// An IntHeap is a min-heap of ints.
// type IntHeap []int

// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *IntHeap) Push(x any) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

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
		// On benchmarking, heap is slower than the list sort, for the example atleast
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
	sort.Ints(leftList)
	sort.Ints(rightList)

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
