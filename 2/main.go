package main

import (
	"bufio"
	"fmt"
	"os"
)

func line(id string) (bool, bool) {
	return true, true
}

func main() {
	f, err := os.Open("./input/input.txt")
	// f, err := os.Open("./input/test.txt")
	if err != nil {
		fmt.Println("there is error")
		return
	}

	scanner := bufio.NewScanner(f)
	two := 0
	three := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := make(map[rune]int)
		hasTwo := false
		hasThree := false

		for _, char := range line {
			chars[char]++
		}
		for _, num := range chars {
			if num == 2 && !hasTwo {
				hasTwo = true
			}
			if num == 3 && !hasThree {
				hasThree = true
			}
		}

		if hasTwo {
			two++
		}
		if hasThree {
			three++
		}
	}
	fmt.Println("Checksum:", two*three)
}
