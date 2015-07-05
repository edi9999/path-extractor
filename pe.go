package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := os.Stdin
	if scanner := bufio.NewScanner(stdin); scanner != nil {
		for scanner.Scan() {
			matches := getAllMatches(scanner.Text())
			for _, match := range matches {
				fmt.Println(match)
			}
		}
	}
}
