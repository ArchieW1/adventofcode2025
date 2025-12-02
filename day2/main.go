package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
		count, err = solution(file, isRepeatedDigits)
	case 2:
		count, err = solution(file, isRepeatedDigits2)
	default:
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Failed counting invalid ids with err: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Total invalid ids is: %d\n", count)
}

func ScanCommaDeliminated(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	nextCommaIndex := bytes.IndexByte(data, ',')
	if nextCommaIndex >= 0 {
		return nextCommaIndex + 1, data[:nextCommaIndex], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

type InvalidNumFunc func(n int) bool

func solution(r io.Reader, b InvalidNumFunc) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanCommaDeliminated)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		indexDash := strings.IndexByte(line, '-')
		lower, err := strconv.Atoi(line[:indexDash])
		if err != nil {
			return 0, err
		}
		upper, err := strconv.Atoi(line[indexDash+1:])
		if err != nil {
			return 0, err
		}

		for i := lower; i <= upper; i++ {
			if b(i) {
				count += i
			}
		}
	}
	return count, scanner.Err()
}

func isRepeatedDigits(n int) bool {
	str := strconv.Itoa(n)
	if len(str)%2 == 1 {
		return false
	}

	m := len(str) / 2
	for i := 0; i < m; i++ {
		if str[i] != str[m+i] {
			return false
		}
	}

	return true
}

func isRepeatedDigits2(n int) bool {
	str := strconv.Itoa(n)

	for numOfP := 2; numOfP <= len(str); numOfP++ {
		if len(str)%numOfP != 0 {
			continue
		}

		pLen := len(str) / numOfP
		good := true
		for i := 0; i < pLen; i++ {
			for j := 1; j < numOfP; j++ {
				if str[i] != str[pLen*j+i] {
					good = false
					break
				}
			}
			if !good {
				break
			}
		}
		if good {
			return true
		}
	}

	return false
}
