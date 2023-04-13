package parsers

func ParsePostgresArray(initial string) []string {
	var lines []string
	var curLine string

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
