package main

import "strings"
import "regexp"

func isGitRange (input string) bool {
    r := regexp.MustCompile("[0-9a-f]{3,}\\.\\.[0-9a-f]{3,}")
    return r.Match([]byte(input))
}

func isGitPath (input string) bool {
    r := regexp.MustCompile("^[ab]/")
    return r.Match([]byte(input))
}

func isDate (input string) bool {
    r := regexp.MustCompile("^[0-9]+/[0-9]+/[0-9]+")
    return r.Match([]byte(input))
}

func isGitInstruction (input string) bool {
    r := regexp.MustCompile("\\.{3,}")
    return r.Match([]byte(input))
}

func replaceGitPath (input string) string {
    r := regexp.MustCompile("^[ab]/(.*)")
    temp := [][]byte{}
    temp = r.FindSubmatch([]byte(input))
    return string(temp[1])
}

func isVersion (input string) bool {
    r := regexp.MustCompile("[0-9x]\\.[0-9x]{1,2}(\\.[0-9x]{1,3})?")
    return r.Match([]byte(input))
}

func containsInvalidString (input string) bool {
    invalidStrings := []string{"(",")","@","origin/","{","}","<",">","$","*"}
    for _,s := range invalidStrings {
        if strings.Contains(input,s) {
            return true
        }
    }
    return false
}
