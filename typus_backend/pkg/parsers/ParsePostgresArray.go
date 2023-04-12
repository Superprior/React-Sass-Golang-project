package parsers

import (
	"strings"
)

func ParsePostgresArray(initial string) (result []string) {
	stripped := initial[1 : len(initial)-1]

	var lines []string
	var curLine string
	var insideLine bool = true

	for i := 1; i < len(stripped); i++ {
		if stripped[i] == '"' && stripped[i-1] != '\\' {
			if insideLine {
				curLine = strings.Replace(curLine, "\\\\t", "    ", -1)
				curLine = strings.Replace(curLine, "\\\\n", "", -1)
				curLine = strings.Replace(curLine, "\\", "", -1)
				lines = append(lines, curLine)
				curLine = ""
				insideLine = false
			} else {
				insideLine = true
			}
		} else {
			if insideLine {
				curLine += string(stripped[i])
			}
		}
	}

	return lines
}
