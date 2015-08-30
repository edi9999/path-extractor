package main

import "github.com/edi9999/path-extractor"

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := os.Stdin
	if scanner := bufio.NewScanner(stdin); scanner != nil {
		for scanner.Scan() {
			matches := pathextractor.GetAllMatches(scanner.Text(), pathextractor.MatchOptions{})
			for _, match := range matches {
				fmt.Println(match)
			}
		}
	}
}
