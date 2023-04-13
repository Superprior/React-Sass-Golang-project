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
