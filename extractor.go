package pathextractor

import "regexp"

func pathExtractor(input string) [][][]byte {
	surroundRegex := "[^][ \\t:'\"]"
	r := regexp.MustCompile("(" + surroundRegex + "*[\\./]" + surroundRegex + "*)")
	temp := [][][]byte{}
	temp = r.FindAllSubmatch([]byte(input), -1)
	return temp
}

func GetAllMatches(input string) []string {
	matches := [][][]byte{}
	result := []string{}
	s := string("")
	matches = pathExtractor(input)
	for _, match := range matches {
		s = string(match[1])
		if isDate(s) || isVersion(s) || isGitRange(s) || isGitInstruction(s) || containsInvalidString(s) || len(s) <= 2 {
			continue
		}
		if isGitPath(s) {
			s = replaceGitPath(s)
		}
		result = append(result, s)
	}
	return result
}
