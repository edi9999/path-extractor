package pathextractor

import "regexp"
import "strings"
import "fmt"

type MatchOptions struct {
	format string
}

func pathExtractor(input string) [][][]byte {
	surroundRegex := "[@~\\-_a-zA-Z/.0-9]*"
	r := regexp.MustCompile("(" + surroundRegex + "[\\./]" + surroundRegex + ")")
	temp := [][][]byte{}
	temp = r.FindAllSubmatch([]byte(input), -1)
	return temp
}

func stripParens(input string) string {
	r := regexp.MustCompile("^\\((.*)\\)$")
	temp := [][]byte{}
	temp = r.FindSubmatch([]byte(input))
	if len(temp) <= 1 {
		return input
	}
	return string(temp[1])
}

func postProcess(input string) string {
	input = stripParens(input)
	return input
}

func GetAllMatches(input string, options MatchOptions) []string {
	matches := [][][]byte{}
	result := []string{}
	s := string("")
	// print(input)
	matches = pathExtractor(input)
	for _, match := range matches {
		s = string(match[1])
		if len(input) >= len(s+"(") && strings.Index(input, s+"(") != -1 {
			continue
		}

		if isEmail(s) || isDate(s) || isVersion(s) || isGitRange(s) || isGitInstruction(s) || endsWithInvalidString(s) || containsInvalidString(s) || len(s) <= 2 {
			continue
		}
		if isGitPath(s) {
			s = replaceGitPath(s)
		}
		s = postProcess(s)
		if options.format == "ackmate" {
			s = fmt.Sprint(s, ":45")
		}
		result = append(result, s)
	}
	return result
}
