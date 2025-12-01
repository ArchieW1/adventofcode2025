package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
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
		count, err = part1(file)
	case 2:
		count, err = part2(file)
	default:
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Failed counting rotations with err: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Password is: %d\n", count)
}

func part1(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	count := 0
	curr := 50
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			return 0, fmt.Errorf("invalid formatting on line: %s", line)
		}

		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}

		switch line[0] {
		case 'L':
			curr = (curr - num) % 100
			if curr < 0 {
				curr += 100
			}
		case 'R':
			curr = (curr + num) % 100
		default:
			return 0, fmt.Errorf("invalid formatting on line: %s", line)
		}

		if curr == 0 {
			count++
		}
	}

	return count, scanner.Err()
}

func part2(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	count := 0
	curr := 50
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			return 0, fmt.Errorf("invalid formatting on line: %s", line)
		}

		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}

		switch line[0] {
		case 'L':
			num = -1 * num
			reversed := (100 - curr) % 100
			count += (reversed - num) / 100
		case 'R':
			count += (curr + num) / 100
		default:
			return 0, fmt.Errorf("invalid formatting on line: %s", line)
		}

		curr = (curr + num) % 100
		if curr < 0 {
			curr += 100
		}
	}

	return count, scanner.Err()
}
