package pathextractor

import "regexp"
import "strings"
import "fmt"

type MatchOptions struct {
	format string
}

func pathExtractor(input string) [][]int {
	surroundRegex := "[@~\\-_a-zA-Z/.0-9]*"
	r := regexp.MustCompile("(" + surroundRegex + "[\\./]" + surroundRegex + ")")
	return r.FindAllSubmatchIndex([]byte(input), -1)
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
	result := []string{}
	candidatePath := string("")
	indexes := pathExtractor(input)
	for _, index := range indexes {
		candidatePath = input[index[0]:index[1]]
		if len(input) >= len(candidatePath+"(") && strings.Index(input, candidatePath+"(") != -1 {
			continue
		}

		if isEmail(candidatePath) || isDate(candidatePath) || isVersion(candidatePath) || isGitRange(candidatePath) || isGitInstruction(candidatePath) || endsWithInvalidString(candidatePath) || containsInvalidString(candidatePath) || len(candidatePath) <= 2 {
			continue
		}
		if isGitPath(candidatePath) {
			candidatePath = replaceGitPath(candidatePath)
		}
		candidatePath = postProcess(candidatePath)
		lineNumber := 45
		columnNumber := 1
		if options.format == "ackmate" {
			candidatePath = fmt.Sprint(candidatePath, ":", lineNumber, ":", columnNumber)
		}
		result = append(result, candidatePath)
	}
	return result
}
