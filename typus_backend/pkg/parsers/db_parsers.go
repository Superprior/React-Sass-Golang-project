package parsers

import (
	"bufio"
	"strings"
)

func RawStringToLinesArray(initial string) []string {
	var result []string

	sc := bufio.NewScanner(strings.NewReader(initial))
	for sc.Scan() {
		result = append(result, sc.Text())
	}
	return result
}

func LinesArrayToPostgresArray(initial []string) string {
	result := "["

	for i, line := range initial {
		result += string("'" + line + "'")
		if i != len(initial)-1 {
			result += ","
		}
	}

	result += "]"
	return result
}

func ParsePostgresArray(initial string) []string {
	var (
		lines   []string
		curLine string
	)

	for i := 0; i < len(initial)-1; i++ {
		if initial[i] == '\\' && initial[i+1] == '\\' {
			lines = append(lines, curLine)
			curLine = ""
			i += 1
		} else {
			curLine += string(initial[i])
		}
	}
	lines = append(lines, curLine+string(initial[len(initial)-1]))
	return lines
}
