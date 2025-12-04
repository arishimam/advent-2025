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

	curr := 50
	count := 0

	for _, turn := range lines {
		direction := turn[0]
		num, err := strconv.Atoi(turn[1:])

		if err != nil {
			fmt.Println("Error converting string to int on ", turn)
			return
		}

		fullTurns := num / 100
		partialTurn := num % 100

		count += fullTurns

		prev := curr
		if direction == 'L' {
			curr -= partialTurn
		} else if direction == 'R' {
			curr += partialTurn
		} else {
			fmt.Println("NEITHER DIRECTION VALID")
		}

		if prev != 0 && (curr < 0 || curr > 100) {
			count++
		}

		if curr < 0 {
			curr = 100 + curr
		}

		if curr >= 100 {
			curr = curr % 100
		}
		if curr == 0 {
			count++
		}

	}

	fmt.Println("Password = ", count)

}
