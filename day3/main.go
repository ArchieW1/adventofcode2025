package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
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
		count, err = solution(file, maxTwoJoltage)
	case 2:
		count, err = solution(file, maxTwelveJoltage)
	default:
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Failed summing joltage with err: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Total output joltage is: %d\n", count)
}

type maxJoltageFunc func(line string) int

func solution(r io.Reader, m maxJoltageFunc) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += m(line)
	}
	return sum, scanner.Err()
}

func maxTwoJoltage(line string) int {
	max := 0
	maxI := 0
	for i := 0; i < len(line)-1; i++ {
		num := int(line[i] - '0')
		if num > max {
			max = num
			maxI = i
		}
	}
	secMax := 0
	for i := maxI + 1; i < len(line); i++ {
		num := int(line[i] - '0')
		if num > secMax {
			secMax = num
		}
	}
	return max*10 + secMax
}

func maxTwelveJoltage(line string) int {
	final := 0
	prevI := -1
	for i := 11; i >= 0; i-- {
		max := 0
		for j := prevI + 1; j < len(line)-i; j++ {
			num := int(line[j] - '0')
			if num > max {
				max = num
				prevI = j
			}
		}
		final += max * int(math.Pow10(i))
	}
	return final
}
