package pathextractor

import (
	"regexp"
	"strings"
)

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func isGitRange(input string) bool {
	r := regexp.MustCompile("[0-9a-f]{3,}\\.\\.[0-9a-f]{3,}")
	return r.Match([]byte(input))
}

func isGitPath(input string) bool {
	r := regexp.MustCompile("^[ab]/")
	return r.Match([]byte(input))
}

func isEmail(input string) bool {
	r := regexp.MustCompile("[a-zA-Z0-9._%+-]+@(?:[a-zA-Z0-9-]+.)+([a-zA-Z]{2,4})")
	result := r.FindSubmatch([]byte(input))
	if result == nil {
		return false
	}
	fileExtensions := []string{"png", "bmp", "jpeg"}
	return !stringInSlice(string(result[1]), fileExtensions)
}

func isDate(input string) bool {
	r := regexp.MustCompile("^[0-9]+/[0-9]+/[0-9]+")
	return r.Match([]byte(input))
}

func isGitInstruction(input string) bool {
	r := regexp.MustCompile("\\.{3,}")
	return r.Match([]byte(input))
}

func replaceGitPath(input string) string {
	r := regexp.MustCompile("^[ab]/(.*)")
	temp := [][]byte{}
	temp = r.FindSubmatch([]byte(input))
	return string(temp[1])
}

func isVersion(input string) bool {
	r := regexp.MustCompile("^v?[0-9x]{1,3}\\.[0-9x]{1,3}(\\.[0-9x]{1,3})?$")
	return r.Match([]byte(input))
}

func isSpace(input string) bool {
	r := regexp.MustCompile("^[0-9]*\\.[0-9]+[MGK]$")
	return r.Match([]byte(input))
}

func startsWithInvalidString(input string) bool {
	invalidBeginnings := []string{"Error/", "Object.", "Array."}
	for _, s := range invalidBeginnings {
		if strings.Index(input, s) == 0 {
			return true
		}
	}
	return false
}

func endsWithInvalidString(input string) bool {
	invalidEndings := []string{"."}
	for _, s := range invalidEndings {
		if strings.LastIndex(input, s) == len(input)-len(s) {
			return true
		}
	}
	return false
}

func containsInvalidString(input string) bool {
	invalidStrings := []string{"//", "()", "and/or", "remotes/", "origin/", "{", "}", "<", ">", "$", "*", "this."}
	for _, s := range invalidStrings {
		if strings.Contains(input, s) {
			return true
		}
	}
	return false
}
