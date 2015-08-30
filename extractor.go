package pathextractor

import "regexp"

func pathExtractor(input string) [][][]byte {
	surroundRegex := "[^][ \\t:'\"]*"
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

func GetAllMatches(input string) []string {
	matches := [][][]byte{}
	result := []string{}
	s := string("")
	matches = pathExtractor(input)
	for _, match := range matches {
		s = string(match[1])
		if isEmail(s) || isDate(s) || isVersion(s) || isGitRange(s) || isGitInstruction(s) || endsWithInvalidString(s) || containsInvalidString(s) || len(s) <= 2 {
			continue
		}
		if isGitPath(s) {
			s = replaceGitPath(s)
		}
		result = append(result, postProcess(s))
	}
	return result
}
