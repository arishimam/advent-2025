package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func handleInput() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func main() {
	lines, err := handleInput()
	if err != nil {
		panic(err)
	}

	/*
		tester := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
	*/

	curr := 50
	count := 0

	for _, turn := range lines {
		num, err := strconv.Atoi(turn[1:])

		if err != nil {
			fmt.Println("Error converting string to int on ", turn)
			return
		}

		prev := curr
		if turn[0] == 'L' {
			curr -= num
		} else {
			curr += num
		}
		if curr != 0 {
			if (prev > 0 && curr < 0) || (prev < 0 && curr > 0) {
				count++
			}

		}

		val := curr / 100
		if val != 0 {
			if val < 0 {
				val *= -1
			}

			count += val

		}

		curr = curr % 100

		if curr == 0 {
			count++
		}

		if curr < 0 {
			curr = 100 + curr
		}

	}

	fmt.Println("Password = ", count)

}
