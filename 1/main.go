package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input/input.txt")
	if err != nil {
		fmt.Println("there is error")
		return
	}

	scanner := bufio.NewScanner(f)
	frequency := 0
	for scanner.Scan() {
		line := scanner.Text()
		offset, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("There was an error")
		}
		frequency += offset
    }
    
    fmt.Println("Final frequency:", frequency)
}
