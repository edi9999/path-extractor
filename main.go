package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	uniquePtr := flag.Bool("u", false, "unique")
	flag.Parse()

	stdin := os.Stdin
	if scanner := bufio.NewScanner(stdin); scanner != nil {
		keys := make(map[string]bool)
		for scanner.Scan() {
			matches := GetAllMatches(scanner.Text(), "ackmate")
			if *uniquePtr == true {
				newMatches := []string{}
				for _, entry := range matches {
					if _, value := keys[entry]; !value {
						keys[entry] = true
						newMatches = append(newMatches, entry)
					}
				}
				matches = newMatches
			}
			for _, match := range matches {
				fmt.Println(match)
			}
		}
	}
}
