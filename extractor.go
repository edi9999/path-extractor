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

func GetAllMatches(input string, format string) []string {
	options := MatchOptions{format: format}
	result := []string{}
	candidatePath := string("")
	restOfLine := string("")
	indexes := pathExtractor(input)
	for _, index := range indexes {
		candidatePath = input[index[0]:index[1]]
		if len(input) >= len(candidatePath+"(") && strings.Index(input, candidatePath+"(") != -1 {
			continue
		}

		if isIp(candidatePath) || isEmail(candidatePath) || isDate(candidatePath) || isVersion(candidatePath) || isGitRange(candidatePath) || isGitInstruction(candidatePath) || startsWithInvalidString(candidatePath) || endsWithInvalidString(candidatePath) || containsInvalidString(candidatePath) || len(candidatePath) <= 2 || isSpace(candidatePath) {
			continue
		}
		if isGitPath(candidatePath) {
			candidatePath = replaceGitPath(candidatePath)
		}
		candidatePath = postProcess(candidatePath)
		if options.format == "ackmate" {
			restOfLine = input[index[1]:]
			cursorPos := getCursorPosition(restOfLine)
			candidatePath = fmt.Sprint(candidatePath, cursorPos)
		}
		result = append(result, candidatePath)
	}
	return result
}

func getCursorPosition(input string) string {
	r := regexp.MustCompile("^(:[0-9]+(:[0-9]+)?)")
	temp := [][]byte{}
	temp = r.FindSubmatch([]byte(input))
	if len(temp) <= 1 {
		return ""
	}
	return string(temp[1])
}
