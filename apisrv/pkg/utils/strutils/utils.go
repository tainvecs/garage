package strutils

import "regexp"

// StringSlicesXOR returns the elements Exclusive OR
// (in `ss1` that are not in `ss2` or `ss2` that are not in `ss1`)
func StringSlicesXOR(ss1, ss2 []string) []string {

	memDict := make(map[string]bool, 0)
	resSlice := make([]string, 0)

	// add ss2 elements to memDict
	for _, s := range ss2 {
		memDict[s] = true
	}

	// check ss1 elements
	for _, s := range ss1 {

		if _, exist := memDict[s]; exist {
			memDict[s] = false
			continue
		}
		resSlice = append(resSlice, s)
	}

	// check ss2 elements
	for s, notOverlap := range memDict {
		if notOverlap {
			resSlice = append(resSlice, s)
		}
	}

	return resSlice
}

// TrimAllIndent trim all the indent including space or line break of a string
func TrimAllIndent(inStr string) string {
	rgp := regexp.MustCompile(`\s`)
	outStr := rgp.ReplaceAllString(inStr, "")
	return outStr
}
