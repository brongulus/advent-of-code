package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var COLS, ROWS int
var answer = 0
var answer2 = 0

func main() {
	file, err := os.ReadFile("./in")
	if err != nil {
		log.Fatal(err)
	}
	grid := strings.Split(string(file), "\n")
	COLS = len(grid[0])
	ROWS = COLS

	for x, row := range grid {
		for y, char := range row {
			if char == 'X' {
				findp1(grid, x, y)
			}
			if char == 'A' {
				findp2(grid, x, y)
			}
		}
	}

	fmt.Println(answer)
	fmt.Println(answer2)
}

func findp2(grid []string, x, y int) {
	dx := [4]int{1, -1, -1, 1}
	dy := [4]int{1, 1, -1, -1}
	xMasMap := [4][4]byte{
		{'S', 'S', 'M', 'M'},
		{'M', 'M', 'S', 'S'},
		{'S', 'M', 'M', 'S'},
		{'M', 'S', 'S', 'M'},
	}
	for check := range len(xMasMap) {
		inside := true
		for i := range len(dx) {
			if inGrid(x+dx[i], y+dy[i]) && grid[x+dx[i]][y+dy[i]] == xMasMap[check][i] {
				continue
			} else {
				inside = false
				break
			}
		}
		if inside {
			answer2++
		}
	}
}

func findp1(grid []string, x, y int) {
	dx := [8]int{1, -1, 0, 0, 1, -1, -1, 1}
	dy := [8]int{0, 0, 1, -1, 1, 1, -1, -1}

	for i := range len(dx) {
		inside := true
		for dletter, letter := range [4]byte{'X', 'M', 'A', 'S'} {
			xCoord := x + dx[i]*dletter
			yCoord := y + dy[i]*dletter
			if inGrid(xCoord, yCoord) && (grid[xCoord][yCoord] == letter) {
				continue
			} else {
				inside = false
				break
			}
		}
		if inside {
			answer++
		}
	}
}

func inGrid(x, y int) bool {
	return x >= 0 && x < ROWS && y >= 0 && y < COLS
}
