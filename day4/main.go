package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	path := flag.String("i", "./input", "Input file path")
	part := flag.Int("p", 2, "Advent of code part 1 or 2")
	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		fmt.Printf("Failed to get input with err: %v", err)
		flag.Usage()
		os.Exit(1)
	}
	defer file.Close()

	var count int
	switch *part {
	case 1:
		count, err = solution(file, adjToiletPaper)
	case 2:
		count, err = solution(file, repeatedAdjToiletPaper)
	default:
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Failed counting paper with err: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Total paper is: %d\n", count)
}

type accessibleToiletPaperFunc func(grid [][]byte) int

func solution(r io.Reader, m accessibleToiletPaperFunc) (int, error) {
	scanner := bufio.NewScanner(r)

	grid := make([][]byte, 0)
	for scanner.Scan() {
		line := []byte(scanner.Text()) // copy of bytes rather than .Bytes() (internal ptr)
		grid = append(grid, line)
	}
	return m(grid), scanner.Err()
}

func adjToiletPaper(grid [][]byte) int {
	count := 0

	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	// 2d copy
	gridCpy := make([][]byte, len(grid))
	for i, inner := range grid {
		innerCpy := make([]byte, len(inner))
		copy(innerCpy, inner)
		gridCpy[i] = innerCpy
	}

	for i := range gridCpy {
		for j := range gridCpy[i] {
			if gridCpy[i][j] != '@' {
				continue
			}

			adj := 0
			for _, direction := range directions {
				dx, dy := direction[0], direction[1]
				if 0 <= i+dx && i+dx < len(gridCpy) && 0 <= j+dy && j+dy < len(gridCpy[i+dx]) {
					if gridCpy[i+dx][j+dy] == '@' {
						adj++
					}
				}
			}

			if adj < 4 {
				count++
				grid[i][j] = '.' //remove for p2
			}
		}
	}

	return count
}

func repeatedAdjToiletPaper(grid [][]byte) int {
	count := 0
	removed := -1
	for removed != 0 {
		removed = adjToiletPaper(grid)
		count += removed
	}
	return count
}
